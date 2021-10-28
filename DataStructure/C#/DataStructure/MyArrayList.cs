namespace DataStructure
{
    public class MyArrayList<T>
    {
        private int Counter { get; set; }

        private T[] InsideArray { get; set; }

        public MyArrayList()
        {
            InsideArray = new T[5];
        }

        public T Get(int index)
        {
            return InsideArray[index];
        }

        public void Add(T value)
        {
            CheckIfNeedDoubleArray();
            InsideArray[Counter] = value;
            Counter++;
        }

        public void Insert(int index, T value)
        {
            CheckIfNeedDoubleArray();
            for (int i = Counter; i > index; i--)
            {
                InsideArray[i] = InsideArray[i - 1];
            }
            InsideArray[index] = value;
            Counter++;
        }

        public void Replace(int index, T value)
        {
            InsideArray[index] = value;
        }

        public void Clear()
        {
            InsideArray = new T[10];
        }

        public bool Contains(T value)
        {
            foreach (var item in InsideArray)
            {
                if (item.Equals(value))
                {
                    return true;
                }
            }
            return false;
        }

        public void CopyTo(MyArrayList<T> array, int index)
        {
            for (int i = 0; i < InsideArray.Length; i++)
            {
                array.Replace(index + 1, InsideArray[i]);
            }
        }

        public int IndexOf(T value)
        {
            for (int i = 0; i < InsideArray.Length; i++)
            {
                if (InsideArray[i].Equals(value))
                {
                    return i;
                }
            }
            return -1;
        }

        public void Remove(T value)
        {
            var removeCount = 0;
            for (int i = 0; i < InsideArray.Length; i++)
            {
                if (InsideArray[i].Equals(value))
                {
                    removeCount++;
                    for (int j = i; j < InsideArray.Length; j++)
                    {
                        InsideArray[j] = InsideArray[j + 1];
                    }
                    i--;
                }
            }
            Counter -= removeCount;
        }

        public void RemoveAt(int index)
        {
            for (int j = index; j < InsideArray.Length; j++)
            {
                InsideArray[j] = InsideArray[j + 1];
            }
        }

        private void CheckIfNeedDoubleArray()
        {
            if (Counter + 1 == InsideArray.Length)
            {
                var tmp = new T[InsideArray.Length * 2];
                for (int i = 0; i < InsideArray.Length; i++)
                {
                    tmp[i] = InsideArray[i];
                }
                InsideArray = tmp;
            }
        }
    }
}

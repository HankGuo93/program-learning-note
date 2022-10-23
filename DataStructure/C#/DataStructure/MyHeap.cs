using System;

namespace DataStructure
{
    public class MyHeap
    {
        public int[] ArrayData { get; set; }

        public int Counter;

        public MyHeap(int lenth = 100)
        {
            ArrayData = new int[lenth];
        }

        public void Insert(int value)
        {
            ArrayData[Counter] = value;
            Counter++;
            Swin(Counter - 1);
        }

        public int Extract()
        {
            var maxValue = ArrayData[0];
            ArrayData[0] = 0;
            Counter--;
            MaxHeapify(0);
            return maxValue;
        }

        public int Peek()
        {
            var maxValue = ArrayData[0];
            return maxValue;
        }

        public int[] Sort()
        {
            throw new NotImplementedException();
        }

        private void Swin(int index)
        {
            while (index > 0 && ArrayData[(index - 1) / 2] < ArrayData[index])
            {
                var temp = ArrayData[(index - 1) / 2];
                ArrayData[(index - 1) / 2] = ArrayData[index];
                ArrayData[index] = temp;
                index = (index - 1) / 2;
            }
        }

        private void MaxHeapify(int index)
        {
            var left = 2 * index + 1;
            var right = 2 * index + 2;
            var largetNode = index;

            if (left <= Counter && ArrayData[left] > ArrayData[largetNode])
            {
                largetNode = left;
            }

            if (right <= Counter && ArrayData[right] > ArrayData[largetNode])
            {
                largetNode = right;
            }

            if (largetNode != index)
            {
                var temp = ArrayData[index];
                ArrayData[index] = ArrayData[largetNode];
                ArrayData[largetNode] = temp;
                MaxHeapify(largetNode);
            }
        }
    }
}

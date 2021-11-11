using System;

namespace DataStructure
{
    public class MyQueue<T>
    {
        private MyLinkedList<T> InsideLinkedList { get; set; } = new MyLinkedList<T>();

        public T Dequeue()
        {
            var node = InsideLinkedList.RemoveFirst();
            return node.Value;
        }

        public void Enqueue(T value)
        {
            InsideLinkedList.AddLast(value);
        }

        public T Peek()
        {
            return InsideLinkedList.First.Value;
        }

        public T[] ToArray()
        {
            var tmpArr = new T[InsideLinkedList.Count];
            MyNode<T> tmpNode = InsideLinkedList.First;
            for (int i = 0; i < InsideLinkedList.Count; i++)
            {
                tmpArr[i] = tmpNode.Value;
                tmpNode = tmpNode.Next;
            }
            return tmpArr;
        }
    }
}

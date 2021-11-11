using System;

namespace DataStructure
{
    public class MyStack<T>
    {
        private MyLinkedList<T> InsideLinkedList { get; set; } = new MyLinkedList<T>();

        public T Pop()
        {
            var node = InsideLinkedList.RemoveLast();
            return node.Value;
        }

        public void Push(T value)
        {
            InsideLinkedList.AddLast(value);
        }

        public T Peek()
        {
            return InsideLinkedList.Last.Value;
        }

        public T[] ToArray()
        {
            InsideLinkedList.Reverse();
            var tmpArr = new T[InsideLinkedList.Count];
            MyNode<T> tmpNode = InsideLinkedList.First;
            for (int i = 0; i < InsideLinkedList.Count; i++)
            {
                tmpArr[i] = tmpNode.Value;
                tmpNode = tmpNode.Next;
            }
            InsideLinkedList.Reverse();
            return tmpArr;
        }
    }
}

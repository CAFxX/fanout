package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 int64, p1 uint16) int64 {
    var v6 int = -1
    v7 := []int{}
    fmt.Println(p0, v6)
    var v8 int8 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int8, p1 int8, p2 uint) int64 {
    if (0 == 16) {
        fmt.Println(p1, p1)
        if true {
            var v9 string = ("a" + (("" + ("\n\t" + ("foo bar" + (("\n\t" + ("hello" + "\n\t")) + "\n\t")))) + ((("hello" + ("" + "hello")) + "") + "a")))
            p0 -= p1
            _ = v9
        }
    } else {
        v10 := GenS1{F0: 2147483647, F1: 2147483647, F2: 218}
        _ = v10
    }
    p0 = 42
    var v11 int8 = 127
    v12 := []int{}
    if (0.1 >= 1e10) {
        fmt.Println(p2, p1)
        if ((3.14159 > -2.25) || (!true)) {
            _ = len(v12)
        } else {
            if true {
                var v13 uint16 = 2
                for fi14 := 0; fi14 < 14; fi14++ {
                    v15 := GenS1{F0: -1, F1: 2147483648, F2: 127}
                    v13 -= 968
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 298
}

func main() {
    var v1 int16 = 32767
    v2 := []int{32767, 1, -1, 823}
    v3 := make(map[uint32]string)
    var v4 int8 = 1
    v5 := []uint32{2147483648, 42, 4294967295, 2147483648}
    var v16 uint = 535
    v16 = (v16 ^ v16)
    var v17 uint8 = 128
    v18 := GenS1{F0: 2147483647, F1: 65535, F2: (v17 * v17)}
    v17 += (v17 | 200)
    v19 := GenS2{F0: "hello"}
    if (v17 > (v17 / ((v17) | 1))) {
        fmt.Println(v4)
        var v20 uint64 = 2
        var v21 uint64 = (^(v20))
        _ = v20
        _ = v21
    }
    fmt.Println(v1, v16, v1)
    v17 += 154
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}

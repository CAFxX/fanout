package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint {
    switch -215 {
    case -2147483648:
        v4 := GenS1{F0: "\n\t"}
        _ = v4
    case 665:
        if ((-2.25 + 1e10) < -2.25) {
            v5 := []int{0, 4294967295, 4294967295}
            if true {
                v6 := GenS1{F0: (((("" + ("" + "a")) + "\n\t") + (("" + ("foo bar" + ("" + ("a" + (("foo bar" + (("foo bar" + "a") + "")) + "\n\t"))))) + "foo bar")) + ("" + "foo bar"))}
                _ = v6
            } else {
                var v7 float64 = (0.0 - 1e10)
                _ = v7
            }
            _ = v5
        }
    }
    v8 := []int{688, 127, -958, 32767}
    _ = v8
    return 65535
}
func genFn1(p0 string) int8 {
    v9 := []int{32768}
    v10 := GenS1{F0: "a"}
    _ = len(v9)
    _ = p0
    _ = v10
    _ = v9
    return 8
}

func main() {
    var v1 uint32 = 0
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    v11 := []int{2147483647, 32767}
    var v12 int16 = -898
    _ = len(v11)
    v2 = 0
    if false {
        if false {
            v13 := GenS1{F0: ""}
            v12 -= v12
            _ = v13
        } else {
            delete(v3, (v1 - (v1 >> ((42) & 63))))
        }
        v1 -= 0
        fmt.Println(v2)
    } else {
        var v14 int16 = (v12 * -32768)
        var v15 string = (("hello" + "") + "a")
        _ = v14
        _ = v15
    }
    var v16 uint32 = v1
    delete(v3, 42)
    for fi17 := 0; fi17 < 1; fi17++ {
        v12 += -1
        _ = len(v3)
        _ = fi17
    }
    fmt.Println(v2)
    v2 += (v2 - (v2 % ((42) | 1)))
    v18 := GenS1{F0: ("hello" + (("" + "foo bar") + "\n\t"))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}

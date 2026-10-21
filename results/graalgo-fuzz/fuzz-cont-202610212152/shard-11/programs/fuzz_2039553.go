package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint16) int64 {
    p0 += 970
    var v5 string = "a"
    v6 := []int{-2147483648, 0}
    v7 := GenS1{F0: true}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint {
    v8 := []int{}
    var v9 uint32 = 0
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = v8
    _ = v9
    return 841
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 544
    var v3 string = "hello"
    var v4 int16 = 32767
    var v10 int32 = -2147483648
    if false {
        for fi11 := 0; fi11 < 18; fi11++ {
            v3 = "foo bar"
            v3 = (v3 + (v3 + "\n\t"))
            _ = fi11
        }
        v4 = v4
    }
    var v12 uint32 = 65535
    fmt.Println(v3, v2)
    for fi13 := 0; fi13 < 3; fi13++ {
        if false {
            delete(v1, (^(v12)))
            var v14 uint16 = 255
            var v15 uint = 0
            _ = v14
            _ = v15
        }
        var v16 int16 = (v4 << (((v2 >> (((^(v2))) & 63))) & 63))
        v17 := GenS1{F0: true}
        _ = fi13
        _ = v16
        _ = v17
    }
    v10 = v10
    fmt.Println(v4)
    fmt.Println(v2, v12)
    if (-2.25 != -2.25) {
        for fi18 := 0; fi18 < 7; fi18++ {
            fmt.Println(v3, v12)
            _ = fi18
        }
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 int64 = 42
            _ = fi19
            _ = v20
        }
        if ((^(v10)) > v10) {
            switch 91 {
            case 4294967295:
                v21 := []int{}
                _ = v21
            case 9223372036854775807:
                v10 -= v10
            case 32767:
                if (42 != 42) {
                    var v22 uint32 = (v12 >> (((v2 >> ((v2) & 63))) & 63))
                    _ = v22
                }
            }
            v4 -= v4
        }
    }
    fmt.Println(v3, v2)
    v23 := []int{int(v2)}
    v24 := GenS1{F0: true}
    fmt.Println(v2)
    v3 = ((v3 + (("hello" + "") + v3)) + (("" + ("\n\t" + v3)) + ("foo bar" + v3)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}

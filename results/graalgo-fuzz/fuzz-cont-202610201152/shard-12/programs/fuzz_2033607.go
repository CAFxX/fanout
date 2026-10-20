package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint) uint8 {
    fmt.Println(p0)
    v4 := []int{-1, 42}
    switch (("\n\t" + "a") + ("hello" + "")) {
    case "\n\t":
        for fi5 := 0; fi5 < 4; fi5++ {
            if false {
                var v6 int16 = 42
                p0 -= p0
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v6
            } else {
                var v7 float32 = (float32(1e10) + 0.0)
                fmt.Println(fi5, p0, v7)
                _ = v7
            }
            _ = fi5
        }
    case "foo bar":
        fmt.Println(p0)
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 bool = true
            v9 = true
            _ = fi8
            _ = v9
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    _ = p0
    _ = v4
    return uint8(p0)
}
func genFn1(p0 int16, p1 int32) bool {
    if false {
        var v10 uint64 = 777
        v11 := []int{4294967295, -702, -2147483648, -917}
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: "hello", F1: ""}
    _ = p0
    _ = p1
    _ = v12
    return false
}

func main() {
    var v1 float32 = (0.1 - 1.5)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v13 := []int{42, 42}
    if (("\n\t" == (("hello" + ("" + ("" + ""))) + "hello")) && (true && true)) {
        v2[909] = 2147483647
        _ = len(v3)
        switch ("\n\t" + "") {
        case "\n\t":
            fmt.Println(v1)
        case "a":
            v1 += 0.0
        }
    } else {
        if true {
            v14 := GenS1{F0: (("" + (("\n\t" + "\n\t") + "")) + ("a" + "")), F1: "foo bar"}
            fmt.Println(v1)
            _ = v14
        }
    }
    if false {
        v1 -= ((0.0 + v1) + float32((3.14159 / (-2.25 + ((1.5 * 1.5) * ((1e10 - ((1.5 * 3.14159) / 3.14159)) + (float64(0.0) / 1e10)))))))
        v15 := GenS2{F0: 32767}
        _ = v15
    }
    var v16 uint32 = 2
    if true {
        var v17 uint64 = 734
        _ = v17
    } else {
        if true {
            fmt.Println(v1, v1)
            v16 -= v16
        }
        v16 += (v16 - (v16 << ((4294967295) & 63)))
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    v2[128] = (v16 >> ((844) & 63))
    var v18 int32 = 0
    var v19 float32 = 0.1
    v16 = v16
    v18 -= (v18 & (-(v18)))
    v20 := GenS2{F0: 0}
    v2[2147483647] = (^(v16))
    if (9223372036854775807 == -9223372036854775808) {
        v21 := GenS1{F0: "", F1: "\n\t"}
        _ = v21
    }
    if (136 >= 0) {
        for fi22 := 0; fi22 < 2; fi22++ {
            v18 += 2147483647
            v16 = (v16 + (v16 | v16))
            _ = fi22
        }
    }
    var v23 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}

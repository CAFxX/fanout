package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int; F2 bool }
type GenS3 struct { F0 int }
func genFn0(p0 string) int8 {
    v5 := []int{818, -1, 9223372036854775807, 32767}
    var v6 uint = 65535
    var v7 bool = true
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []string{"foo bar", "a", "a", ("foo bar" + ((("" + "a") + "foo bar") + "foo bar"))}
    v2 := make(map[string]int)
    var v3 float64 = (float64(0.0) + 1.5)
    var v4 int32 = 1
    if len(v1) > 0 {
        v1[0] = ""
    }
    var v8 int64 = 191
    v8 += 9223372036854775807
    v4 = -1
    var v9 uint32 = 1
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(v4, v9)
        fmt.Println(fi10, v3, v8)
        var v11 bool = false
        _ = fi10
        _ = v11
    }
    v8 = 254
    v12 := GenS1{F0: (!true)}
    fmt.Println(v3, v4, v4)
    for fi13 := 0; fi13 < 1; fi13++ {
        delete(v2, (("hello" + "\n\t") + "hello"))
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS2{F0: (v8 << ((1) & 63)), F1: fi13, F2: true}
            if true {
                fmt.Println(fi13, v4, v3)
            }
            v16 := GenS2{F0: v8, F1: fi14, F2: (true && (32767 != 1000))}
            _ = fi14
            _ = v15
            _ = v16
        }
        if (328 != 18446744073709551615) {
            v9 = 781
            delete(v2, "")
            switch (fi13 * (fi13 | fi13)) {
            case -2147483648:
                v9 -= 0
            case 4294967295:
                for fi17 := 0; fi17 < 15; fi17++ {
                    v2["\n\t"] = -9223372036854775808
                    fmt.Println(v9, v9, fi13)
                    _ = fi17
                }
            default:
                v18 := GenS3{F0: fi13}
                _ = v18
            }
        } else {
            var v19 uint64 = 542
            _ = v19
        }
        _ = fi13
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        v8 -= v8
        _ = fi20
    }
    var v21 uint = 255
    fmt.Println(v8, v9)
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 string = ("foo bar" + "a")
        v8 -= (v8 & (^(v8)))
        v24 := GenS3{F0: (fi22 ^ fi22)}
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int32 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}

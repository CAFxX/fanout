package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint64; F2 uint8 }
func genFn0(p0 uint8, p1 int64) uint16 {
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = (p0 + 181)
        v7 := []int{fi6, fi6, -1, fi6}
        fmt.Println(p1)
        _ = fi6
        _ = v7
    }
    if ("" > ("\n\t" + "foo bar")) {
        var v8 int16 = 1
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint8 = 255
            p1 += p1
            v11 := []int{(fi9 + (^(fi9)))}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint16 = 32767
        _ = v12
        _ = v8
    } else {
        if true {
            if false {
                var v13 float64 = 0.0
                _ = v13
            }
            for fi14 := 0; fi14 < 9; fi14++ {
                p0 += (^(p0))
                if (false || false) {
                    var v15 int16 = 401
                    var v16 int = -9223372036854775808
                    _ = v15
                    _ = v16
                } else {
                    var v17 int32 = 38
                    p0 = p0
                    _ = v17
                }
                fmt.Println(fi14, fi14)
                _ = fi14
            }
            p0 = (p0 << ((42) & 63))
        } else {
            var v18 int64 = p1
            _ = v18
        }
        fmt.Println(p0, p0)
        var v19 string = ("foo bar" + "")
        _ = v19
    }
    _ = p0
    _ = p1
    return 562
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := []string{("\n\t" + ("a" + ("hello" + "a"))), "\n\t"}
    var v4 int = -153
    var v5 uint8 = 255
    var v20 string = "a"
    var v21 uint8 = 1
    if true {
        v22 := GenS1{F0: 9223372036854775807}
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 int = (v4 >> ((42) & 63))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v25 := GenS2{F0: "a"}
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = v22
    }
    v4 = (v4 % ((445) | 1))
    v26 := GenS3{F0: 42, F1: 369, F2: 186}
    v20 = (v20 + "a")
    switch (v4 + v4) {
    case 2147483647:
        var v27 uint = 42
        _ = v27
    case -452:
        v20 = (v20 + v20)
    case -2147483648:
        fmt.Println(v4, v4, v5)
    default:
        v28 := GenS3{F0: 478, F1: 2, F2: (v5 & v5)}
        _ = v28
    }
    var v29 int64 = -706
    for fi30 := 0; fi30 < 1; fi30++ {
        switch (v20 + "\n\t") {
        case "":
            var v31 string = v20
            _ = v31
        case "a":
            fmt.Println(v4, v29, fi30)
        default:
            fmt.Println(v4, v20)
        }
        var v32 uint32 = 2147483647
        _ = fi30
        _ = v32
    }
    v3 = append(v3, ((v20 + v20) + ("" + "\n\t")))
    v3 = append(v3, "a")
    fmt.Println(v20, v29)
    v33 := GenS1{F0: -9223372036854775808}
    v4 -= v4
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}

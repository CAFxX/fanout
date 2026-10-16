package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() bool {
    var v5 int8 = 127
    if (!(1 >= 255)) {
        v6 := GenS1{F0: 65535}
        _ = v6
    }
    switch 2147483647 {
    case 949:
        v5 = v5
    case 32767:
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int64 = -1
            var v9 uint16 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v5 = v5
    _ = v5
    return (18446744073709551615 != 398)
}
func genFn1(p0 uint64, p1 uint8, p2 uint16) uint16 {
    if true {
        var v10 uint16 = 255
        _ = v10
    } else {
        p1 += (p1 - 128)
        v11 := []int{}
        _ = v11
    }
    switch ("" + "foo bar") {
    case "a":
        switch 198 {
        case 128:
            p0 = 928
        case 4294967295:
            var v12 uint = 646
            _ = v12
        case 1:
            p2 += p2
        default:
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(p2)
                var v14 uint32 = 729
                var v15 bool = (42 != 1)
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    var v16 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 796
}

func main() {
    var v1 int = 128
    v2 := []string{"", ("" + "\n\t"), ("" + ("" + "a")), ("hello" + "foo bar")}
    var v3 int = v1
    var v4 bool = false
    if (true || (!(!true))) {
        v2 = append(v2, "")
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 int32 = 877
        for fi19 := 0; fi19 < 7; fi19++ {
            if v4 {
                if v4 {
                    fmt.Println(v3, fi17, v18)
                    v1 -= (v3 | fi17)
                    var v20 float64 = 0.0
                    _ = v20
                } else {
                    v1 = v3
                }
            }
            v3 = (v3 & (fi17 + (v3 ^ (v1 >> ((126) & 63)))))
            _ = fi19
        }
        var v21 bool = (94 <= -991)
        _ = fi17
        _ = v18
        _ = v21
    }
    var v22 float64 = (0.0 - 0.0)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 += v3
    v1 -= -2147483648
    var v23 uint32 = 544
    switch "\n\t" {
    case "":
        v23 += v23
    case "a":
        if v4 {
            fmt.Println(v23)
            fmt.Println(v22, v22)
            var v24 int8 = -27
            _ = v24
        }
    default:
        var v25 float64 = v22
        _ = v25
    }
    var v26 int8 = -128
    var v27 int16 = 32767
    var v28 int = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}

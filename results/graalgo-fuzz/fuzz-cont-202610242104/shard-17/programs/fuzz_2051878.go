package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
func genFn0(p0 string) string {
    v6 := GenS1{F0: -1, F1: "a", F2: true}
    v7 := GenS1{F0: 1, F1: ((p0 + p0) + (p0 + (p0 + "hello"))), F2: true}
    _ = p0
    _ = v6
    _ = v7
    return p0
}
func genFn1() int64 {
    var v8 uint8 = 58
    fmt.Println(v8)
    var v9 int16 = -625
    v10 := []int{128, int(v8)}
    if (true && true) {
        if ((v9 >= v9) || true) {
            fmt.Println(v9, v8)
        } else {
            if (!true) {
                var v11 uint64 = 42
                if true {
                    v12 := GenS1{F0: v9, F1: "\n\t", F2: (!true)}
                    _ = v12
                }
                v8 += (v8 * v8)
                _ = v11
            }
        }
    }
    _ = v10
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := []string{"foo bar", (("\n\t" + "") + ("\n\t" + "a")), "", "\n\t"}
    v2 := []int{9223372036854775807, -1}
    v3 := []string{""}
    v4 := make(map[int]uint32)
    v5 := []string{"hello", ("\n\t" + "\n\t"), ("\n\t" + (((("\n\t" + ("a" + ("foo bar" + "a"))) + "\n\t") + "\n\t") + "a")), ("\n\t" + "foo bar")}
    if true {
        v4[4294967295] = 1
        var v13 string = "a"
        _ = v13
    } else {
        v14 := GenS1{F0: 1000, F1: "", F2: (false || true)}
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 int32 = -1
        v16 = v16
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint32 = 2147483648
        for fi19 := 0; fi19 < 1; fi19++ {
            v18 = v18
            fmt.Println(fi19, fi17)
            var v20 int32 = 0
            _ = fi19
            _ = v20
        }
        var v21 float64 = ((1e10 + 3.14159) + 0.0)
        _ = fi17
        _ = v18
        _ = v21
    }
    var v22 int8 = 0
    var v23 int8 = v22
    var v24 uint64 = 613
    v25 := GenS1{F0: 1000, F1: "hello", F2: true}
    var v26 uint32 = 0
    v1 = append(v1, ("" + (("foo bar" + "\n\t") + "")))
    var v27 uint32 = 2
    for fi28 := 0; fi28 < 19; fi28++ {
        switch fi28 {
        case -9223372036854775808:
            fmt.Println(v22)
        default:
            fmt.Println(v24, v27, v26)
        }
        _ = fi28
    }
    v27 = v26
    var v29 float64 = -2.25
    v29 = float64(-2.25)
    if false {
        v23 += 74
        switch 32768 {
        case 42:
            var v30 bool = (1 != 42)
            _ = v30
        }
    }
    v29 -= 0.0
    v31 := GenS1{F0: 723, F1: ("hello" + "\n\t"), F2: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}

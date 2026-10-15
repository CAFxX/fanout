package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int8, p1 int) string {
    p1 += (p1 * p1)
    var v7 int = 32767
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := []int{fi9, 32767, (^(v7)), fi9}
        if true {
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS1{F0: 366}
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                v8 -= (v8 >> ((773) & 63))
                v14 := GenS1{F0: 2}
                v8 = 2
                v15 := GenS2{F0: "a"}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        var v16 uint32 = 2
        _ = fi9
        _ = v10
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1() int8 {
    if (!true) {
        var v17 uint16 = 1
        _ = v17
    }
    var v18 int = 261
    v19 := GenS3{F0: "a"}
    var v20 float64 = float64(0.1)
    var v21 int8 = -91
    _ = v18
    _ = v19
    _ = v20
    _ = v21
    return v21
}

func main() {
    var v1 uint = 65535
    v2 := []int{-9223372036854775808, 4294967295, -9223372036854775808}
    var v3 uint16 = 2
    v4 := []uint32{2147483648, uint32(v1), 0}
    var v5 int = 0
    v6 := []string{(("foo bar" + (("\n\t" + (("hello" + "a") + (("a" + "foo bar") + ""))) + ("foo bar" + (("" + "\n\t") + "a")))) + (("" + "a") + "hello")), "hello"}
    v22 := GenS4{F0: 2147483647}
    switch "" {
    case "a":
        v1 += v1
    case "\n\t":
        v23 := GenS2{F0: "hello"}
        _ = v23
    default:
        v5 -= 42
    }
    var v24 string = "hello"
    _ = len(v2)
    v3 = v3
    var v25 int8 = 100
    var v26 int16 = 743
    v5 += v5
    v2 = append(v2, v5)
    v27 := GenS3{F0: "foo bar"}
    var v28 uint8 = 127
    var v29 string = (v24 + "")
    v24 = v29
    v2 = append(v2, v5)
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := GenS3{F0: v29}
            v28 = 128
            var v33 uint = v1
            _ = len(v6)
            v28 = v28
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 int16 = v26
        if ((v1 >> ((0) & 63)) > v1) {
            v5 = (v5 | v5)
            if true {
                fmt.Println(v3, v29)
                var v35 uint = v1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v35
            } else {
                var v36 int64 = -998
                for fi37 := 0; fi37 < 2; fi37++ {
                    var v38 uint8 = 0
                    _ = fi37
                    _ = v38
                }
                v39 := GenS3{F0: v29}
                _ = v36
                _ = v39
            }
        }
        _ = fi30
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

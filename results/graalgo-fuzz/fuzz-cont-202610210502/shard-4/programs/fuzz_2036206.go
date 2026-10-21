package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool }
func genFn0(p0 string, p1 uint32, p2 uint8) uint32 {
    var v5 uint8 = (p2 << ((196) & 63))
    var v6 int64 = 1
    v5 = (p2 ^ (p2 & p2))
    for fi7 := 0; fi7 < 11; fi7++ {
        p2 = v5
        var v8 float32 = float32(1.5)
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: p1, F1: "a", F2: p2}
            v6 -= v6
            if false {
                v11 := []int{(^(fi7)), (fi7 * 32768), (fi9 - fi7)}
                v8 = ((v8 / (v8 / -2.25)) / (v8 / (v8 - 1e10)))
                var v12 float64 = (float64((1.5 - (v8 / v8))) + (1.5 / 3.14159))
                _ = v11
                _ = v12
            } else {
                p0 = p0
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v13 := GenS2{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v5
    _ = v6
    return (p1 - (^(p1)))
}

func main() {
    var v1 string = ("a" + ("" + ("hello" + "")))
    v2 := make(map[int]int)
    var v3 string = v1
    v4 := []string{v1, "hello"}
    if true {
        v14 := GenS2{F0: 65535}
        fmt.Println(v1, v1)
        v3 = "foo bar"
        _ = v14
    } else {
        delete(v2, 1)
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint8 = 0
            _ = len(v4)
            v16 = v16
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: 2147483648}
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v19 := v2[(fi18 * (fi18 >> ((4294967295) & 63)))]
        v1 = (v1 + v1)
        v3 = v1
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        if (!((1.5 - (0.0 / 0.1)) == -2.25)) {
            v21 := GenS1{F0: genFn0(((v3 + "\n\t") + (v3 + ((v1 + "hello") + "\n\t"))), 820, 127), F1: v1, F2: 42}
            if (65535 > 65535) {
                var v22 bool = (2147483647 >= 455)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v22
            } else {
                _ = len(v2)
            }
            var v23 float64 = 0.1
            _ = v21
            _ = v23
        }
        switch 128 {
        case 42:
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := GenS1{F0: 1, F1: "foo bar", F2: 2}
                _ = fi24
                _ = v25
            }
        default:
            v26 := v2[fi20]
            _ = v26
        }
        fmt.Println(fi20)
        _ = fi20
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v3 = ((v3 + "hello") + v3)
        v3 = v1
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(fi28, v3)
            v29 := GenS3{F0: v1, F1: (!(1 > 572))}
            fmt.Println(fi27)
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    v1 = ("hello" + ((v3 + ((((((v1 + "a") + v3) + (v1 + v1)) + v3) + "a") + "foo bar")) + "a"))
    for fi30 := 0; fi30 < 4; fi30++ {
        delete(v2, fi30)
        v3 = ("" + ("foo bar" + ""))
        for fi31 := 0; fi31 < 6; fi31++ {
            var v32 uint16 = 2
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    v3 = ("hello" + "a")
    var v33 float64 = 0.1
    if (true || (true && true)) {
        _ = len(v4)
        v34 := GenS3{F0: v3, F1: true}
        v35 := v2[-2147483648]
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        var v37 string = (v3 + (((v3 + ((v1 + v3) + "\n\t")) + "\n\t") + v3))
        _ = fi36
        _ = v37
    }
    _ = len(v2)
    v2[42] = 32768
    var v38 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
}

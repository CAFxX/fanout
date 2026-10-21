package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 int8; F1 int32; F2 string }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 uint16; F1 bool }
func genFn0(p0 int64) string {
    var v6 int64 = p0
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (fi7 | fi7) {
        case -2147483648:
            var v8 string = ("a" + (((("foo bar" + "a") + ("a" + "")) + "hello") + "a"))
            _ = v8
        }
        var v9 bool = (!(1 > 2147483647))
        fmt.Println(v9, p0)
        _ = fi7
        _ = v9
    }
    _ = p0
    _ = v6
    return "a"
}
func genFn1(p0 int8) string {
    p0 = 21
    v10 := []int{}
    if ("hello" > "\n\t") {
        for fi11 := 0; fi11 < 1; fi11++ {
            if len(v10) > 0 {
                v10[0] = 127
            }
            var v12 int = (-(fi11))
            v12 = 2147483647
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            p0 += p0
            _ = fi13
        }
    } else {
        p0 -= -111
    }
    if true {
        v14 := GenS1{F0: 127}
        _ = v14
    } else {
        switch -1 {
        case -775:
            v15 := GenS1{F0: 2}
            _ = v15
        case -176:
            v16 := GenS2{F0: "a", F1: (220 < 0), F2: 2147483647}
            _ = v16
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            switch "" {
            case "hello":
                fmt.Println(p0, fi17)
            default:
                var v18 int32 = 562
                _ = v18
            }
            _ = fi17
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(p0, fi19)
            var v20 uint16 = 65535
            if len(v10) > 0 {
                v10[0] = -1
            }
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v10
    return ("" + "\n\t")
}

func main() {
    var v1 uint32 = 70
    var v2 float64 = 0.1
    var v3 int16 = 1000
    var v4 float32 = 3.14159
    v5 := make(map[int]string)
    v21 := GenS3{F0: -128, F1: 1, F2: "a"}
    var v22 int8 = -128
    v23 := GenS1{F0: 2}
    fmt.Println(v2, v4, v4)
    if true {
        switch -9223372036854775808 {
        case 32767:
            v4 -= 1e10
        case 127:
            v2 = v2
        }
    } else {
        for fi24 := 0; fi24 < 8; fi24++ {
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 uint = 2
                _ = fi25
                _ = v26
            }
            v22 += (v22 + (-(v22)))
            _ = fi24
        }
        var v27 float64 = v2
        _ = v27
    }
    switch -9223372036854775808 {
    case 127:
        for fi28 := 0; fi28 < 4; fi28++ {
            if ((fi28 & (-(fi28))) > (fi28 ^ (fi28 << ((565) & 63)))) {
                for fi29 := 0; fi29 < 10; fi29++ {
                    v30 := GenS1{F0: 128}
                    if false {
                        var v31 int64 = 9223372036854775807
                        v31 += v31
                        v22 = (v22 + (v22 % (((v22 / (((v22 / ((98) | 1))) | 1))) | 1)))
                        _ = v31
                    }
                    var v32 string = "a"
                    _ = fi29
                    _ = v30
                    _ = v32
                }
                switch ("a" + "hello") {
                case "a":
                    v33 := []int{32768, fi28, fi28}
                    _ = v33
                }
                var v34 int16 = -1
                _ = v34
            } else {
                fmt.Println(v2, v2)
                var v35 int64 = -9223372036854775808
                _ = v35
            }
            v5[-2147483648] = ""
            if false {
                var v36 int64 = -954
                _ = v36
            }
            _ = fi28
        }
    case 42:
        switch "\n\t" {
        case "a":
            for fi37 := 0; fi37 < 1; fi37++ {
                if false {
                    _ = len(v5)
                }
                _ = fi37
            }
        }
    case -74:
        v38 := GenS4{F0: 1, F1: 42}
        _ = v38
    default:
        for fi39 := 0; fi39 < 6; fi39++ {
            var v40 int8 = (v22 + v22)
            v41 := GenS3{F0: (v22 << ((65535) & 63)), F1: 42, F2: "\n\t"}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    v42 := GenS5{F0: 255, F1: false}
    v43 := []int{-9223372036854775808, 485, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}

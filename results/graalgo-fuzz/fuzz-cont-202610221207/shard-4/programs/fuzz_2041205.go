package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint, p1 int, p2 uint) uint16 {
    p1 = p1
    switch ("foo bar" + "a") {
    case "a":
        v7 := GenS1{F0: 878, F1: "foo bar", F2: 31}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 float32 = (1.5 - 0.0)
    var v3 int16 = 156
    var v4 uint8 = 0
    var v5 uint = 0
    var v6 uint = 2
    var v8 float32 = float32(1e10)
    v9 := GenS1{F0: 1, F1: ("" + ((("" + ("a" + "foo bar")) + ("foo bar" + "")) + ("foo bar" + ((("\n\t" + "a") + "a") + (("a" + "a") + "foo bar"))))), F2: -128}
    for fi10 := 0; fi10 < 5; fi10++ {
        v4 += v4
        v6 = (v5 % (((v5 - (v5 - v6))) | 1))
        _ = fi10
    }
    if (2147483647 < 263) {
        v5 += (v5 / ((v6) | 1))
        fmt.Println(v4, v5, v3)
    } else {
        v8 = v8
    }
    v11 := GenS1{F0: 600, F1: "", F2: 0}
    switch 558 {
    case 142:
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int16 = -630
            _ = fi12
            _ = v13
        }
    }
    if ((v3 != (v3 | v3)) && (uint64(v4) > 1)) {
        switch (("a" + "") + ((("hello" + ("foo bar" + ("" + "\n\t"))) + "a") + ("\n\t" + (("\n\t" + "a") + "foo bar")))) {
        case "\n\t":
            var v14 bool = true
            _ = v14
        case "foo bar":
            var v15 int8 = 42
            _ = v15
        }
        fmt.Println(v8)
        var v16 int16 = 1
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 uint8 = (v4 ^ 0)
        v19 := []int{fi17, 0, fi17}
        _ = fi17
        _ = v18
        _ = v19
    }
    switch ("\n\t" + "") {
    case "\n\t":
        v3 += v3
    case "foo bar":
        if (!false) {
            if (true && false) {
                for fi20 := 0; fi20 < 16; fi20++ {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        v4 = 1
                        v5 += v6
                        v8 = v2
                        _ = fi21
                    }
                    if false {
                        v4 = 0
                    } else {
                        v2 = float32(0.0)
                    }
                    _ = fi20
                }
            }
        } else {
            v22 := []int{0, -113, 32767}
            _ = v22
        }
    case "":
        v23 := GenS2{F0: -830}
        _ = v23
    default:
        var v24 int8 = -128
        _ = v24
    }
    v25 := GenS3{F0: 983, F1: -9223372036854775808}
    var v26 float64 = 0.1
    fmt.Println(v2)
    v27 := GenS3{F0: 1, F1: 0}
    for fi28 := 0; fi28 < 2; fi28++ {
        switch (fi28 * fi28) {
        case 0:
            switch fi28 {
            case -301:
                v3 += v3
            case 803:
                var v29 int = fi28
                _ = v29
            default:
                var v30 int32 = -2147483648
                _ = v30
            }
        case 9223372036854775807:
            for fi31 := 0; fi31 < 7; fi31++ {
                v32 := GenS4{F0: 1}
                v33 := []int{(fi28 % (((fi31 ^ -2147483648)) | 1))}
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        v8 = 0.1
        for fi34 := 0; fi34 < 16; fi34++ {
            v4 += v4
            fmt.Println(v4)
            _ = fi34
        }
        _ = fi28
    }
    if true {
        v35 := GenS3{F0: 0, F1: 373}
        _ = v35
    } else {
        v3 = (v3 & v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}

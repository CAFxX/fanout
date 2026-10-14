package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint16 }
func genFn0() int64 {
    var v4 int64 = 1
    var v5 string = ((("foo bar" + "a") + "hello") + "")
    _ = v4
    _ = v5
    return (v4 % ((v4) | 1))
}
func genFn1(p0 bool, p1 uint) int {
    var v6 bool = ((1.5 < float32((0.1 - -2.25))) && (475 <= 65535))
    if (false || v6) {
        fmt.Println(p1)
    } else {
        v7 := GenS1{F0: 161}
        p0 = (-358 >= 10)
        v8 := GenS2{F0: 214, F1: 1}
        _ = v7
        _ = v8
    }
    switch ("\n\t" + "hello") {
    case "a":
        switch "a" {
        case "foo bar":
            fmt.Println(v6, p0, p0)
        case "hello":
            var v9 bool = true
            _ = v9
        default:
            if (1 != 100) {
                for fi10 := 0; fi10 < 1; fi10++ {
                    fmt.Println(p0, v6)
                    _ = fi10
                }
                if (42 >= 665) {
                    p1 = (^(p1))
                    for fi11 := 0; fi11 < 19; fi11++ {
                        switch ("" + "foo bar") {
                        case "":
                            var v12 uint16 = 65535
                            _ = v12
                        default:
                            v13 := GenS3{F0: -32768, F1: 32768}
                            _ = v13
                        }
                        v14 := GenS4{F0: 32767}
                        switch (fi11 - (^(fi11))) {
                        case 2147483647:
                            var v15 int = fi11
                            _ = v15
                        case 612:
                            p0 = ((32767 <= 1) && p0)
                        }
                        _ = fi11
                        _ = v14
                    }
                } else {
                    v16 := GenS2{F0: 42, F1: 881}
                    _ = v16
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}

func main() {
    var v1 string = "hello"
    var v2 int = 42
    var v3 uint8 = 2
    fmt.Println(v3, v2)
    v3 = 255
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := []int{}
        _ = fi17
        _ = v18
    }
    var v19 float64 = float64(0.1)
    v20 := GenS1{F0: 364}
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := []int{v2, (^(v2)), fi21}
        switch 760 {
        case -167:
            v23 := GenS3{F0: -32768, F1: 128}
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v2, v3, v3)
    var v24 int32 = -1
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int32 = -1
        var v27 int8 = int8(v2)
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 int8 = 1
    var v29 int64 = 42
    if false {
        for fi30 := 0; fi30 < 17; fi30++ {
            if (!true) {
                if (!true) {
                    v28 += (v28 - v28)
                } else {
                    v29 -= (v29 & v29)
                }
            } else {
                v24 += (^(v24))
                fmt.Println(v28, v2, v29)
            }
            v24 = 42
            for fi31 := 0; fi31 < 17; fi31++ {
                fmt.Println(v2, fi30, fi31)
                _ = fi31
            }
            _ = fi30
        }
        var v32 int16 = 0
        var v33 int64 = (v29 ^ v29)
        _ = v32
        _ = v33
    }
    var v34 uint8 = (v3 | v3)
    var v35 uint8 = (v3 % ((v34) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
}

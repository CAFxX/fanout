package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 int64 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 0
    var v3 uint64 = 1
    var v4 uint64 = (^(v3))
    v4 -= (^(v4))
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: 128}
        v7 := v1["\n\t"]
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint64 = v4
    if false {
        fmt.Println(v3, v4)
        v9 := GenS1{F0: uint8(v2)}
        _ = v9
    } else {
        if true {
            v8 = (v3 % ((v4) | 1))
            v8 -= 143
        } else {
            fmt.Println(v4, v4)
        }
        v10 := GenS1{F0: 127}
        _ = v10
    }
    var v11 float32 = ((1.5 + -2.25) + ((0.0 / -2.25) + -2.25))
    fmt.Println(v4)
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 string = "foo bar"
        var v14 int = fi12
        _ = fi12
        _ = v13
        _ = v14
    }
    v1["\n\t"] = "a"
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int8 = -128
        var v17 uint8 = 255
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 128}
    if true {
        v19 := GenS1{F0: 127}
        v20 := []int{127, 831, -9223372036854775808}
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int64 = -9223372036854775808
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    } else {
        if (!true) {
            switch "" {
            case "":
                if (((v11 > v11) && (true || true)) || (!true)) {
                    fmt.Println(v2)
                    fmt.Println(v2, v11)
                }
            case "\n\t":
                var v23 uint = 400
                _ = v23
            default:
                v4 -= (v4 - v3)
            }
        } else {
            v24 := GenS1{F0: 49}
            var v25 bool = (false || true)
            _ = v24
            _ = v25
        }
        if true {
            v26 := []int{}
            _ = v26
        }
        if (766 <= 225) {
            if false {
                v4 += (v4 & (v3 % ((v8) | 1)))
                for fi27 := 0; fi27 < 9; fi27++ {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if (v2 == 2) {
                            var v29 float32 = v11
                            _ = v29
                        }
                        v30 := []int{fi27, (fi28 ^ 0)}
                        _ = fi28
                        _ = v30
                    }
                    v4 = v8
                    var v31 uint16 = v2
                    _ = fi27
                    _ = v31
                }
                var v32 float64 = -2.25
                _ = v32
            } else {
                var v33 int8 = 100
                var v34 uint32 = 852
                _ = v33
                _ = v34
            }
        }
    }
    var v35 uint64 = (v4 / ((v3) | 1))
    switch "\n\t" {
    case "hello":
        if false {
            var v36 uint64 = 1
            if true {
                for fi37 := 0; fi37 < 8; fi37++ {
                    var v38 bool = ((false || true) && false)
                    v3 = (v8 | v36)
                    if (!(!false)) {
                        v39 := []int{4294967295}
                        v1[""] = "foo bar"
                        _ = v39
                    }
                    _ = fi37
                    _ = v38
                }
                v40 := GenS1{F0: 255}
                _ = v40
            }
            for fi41 := 0; fi41 < 18; fi41++ {
                fmt.Println(v8)
                v42 := GenS1{F0: 128}
                v43 := v1[("hello" + (("hello" + "") + "hello"))]
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = v36
        }
    case "foo bar":
        v1["\n\t"] = ("a" + "hello")
    default:
        if false {
            v44 := GenS2{F0: 42}
            for fi45 := 0; fi45 < 2; fi45++ {
                v46 := GenS3{F0: 2, F1: 0, F2: 9223372036854775807}
                v2 += 2
                _ = fi45
                _ = v46
            }
            v8 = v3
            _ = v44
        } else {
            var v47 string = ("foo bar" + "hello")
            var v48 uint64 = v35
            for fi49 := 0; fi49 < 17; fi49++ {
                v50 := GenS4{F0: -2147483648}
                _ = fi49
                _ = v50
            }
            _ = v47
            _ = v48
        }
    }
    v8 += (v35 & 4294967295)
    for fi51 := 0; fi51 < 8; fi51++ {
        var v52 uint16 = v2
        v1["hello"] = "\n\t"
        for fi53 := 0; fi53 < 17; fi53++ {
            v54 := []int{(fi53 << ((18446744073709551615) & 63))}
            var v55 bool = false
            _ = fi53
            _ = v54
            _ = v55
        }
        _ = fi51
        _ = v52
    }
    v8 = (v35 + (v4 << ((18446744073709551615) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v8
}

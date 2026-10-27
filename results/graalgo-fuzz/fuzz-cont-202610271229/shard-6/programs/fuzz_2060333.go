package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64) int8 {
    var v6 float64 = (0.0 + 1e10)
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: 1, F1: (1 == 0), F2: "\n\t"}
            _ = fi7
            _ = v8
        }
        fmt.Println(v6)
        v6 = v6
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 bool = false
    var v2 int = -316
    var v3 int64 = 9223372036854775807
    var v4 uint64 = 172
    var v5 uint16 = 695
    var v9 uint64 = (^(v4))
    v5 = 1
    if v1 {
        v5 -= 42
        fmt.Println(v2, v3)
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(v4, v3)
            _ = fi10
        }
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint8 = 0
            for fi13 := 0; fi13 < 3; fi13++ {
                v12 -= v12
                var v14 int32 = 0
                _ = fi13
                _ = v14
            }
            if (1 >= int8(v9)) {
                var v15 int8 = -128
                _ = v15
            }
            _ = fi11
            _ = v12
        }
    }
    var v16 bool = v1
    fmt.Println(v2)
    var v17 int64 = (v3 & v3)
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := []int{(v2 - 0), (^(v2)), (fi18 >> ((18446744073709551615) & 63)), 4294967295}
        for fi20 := 0; fi20 < 17; fi20++ {
            switch int(v5) {
            case 503:
                v9 -= (^(v9))
            case 42:
                if v1 {
                    if (!v16) {
                        v21 := GenS2{F0: (v17 + (v17 * -480))}
                        if (0 < 32767) {
                            v22 := GenS3{F0: 0}
                            switch fi18 {
                            case -961:
                                var v23 int64 = (v3 & -384)
                                _ = v23
                            }
                            fmt.Println(fi18, v2)
                            _ = v22
                        }
                        _ = v21
                    } else {
                        var v24 uint16 = v5
                        var v25 string = "foo bar"
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = v24
                        _ = v25
                    }
                } else {
                    if false {
                        v4 -= 2
                    }
                }
            }
            v1 = (240 <= 65535)
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    if false {
        v26 := GenS4{F0: 1}
        var v27 int = (^(v2))
        _ = v26
        _ = v27
    } else {
        for fi28 := 0; fi28 < 6; fi28++ {
            v3 = (v17 * (v3 & 42))
            _ = fi28
        }
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        var v30 int64 = v3
        var v31 uint = 624
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}

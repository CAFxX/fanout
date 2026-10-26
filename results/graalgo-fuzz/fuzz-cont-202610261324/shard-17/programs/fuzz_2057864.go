package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool) uint16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        p0 = p0
        var v6 int64 = 1
        v7 := GenS1{F0: 1, F1: 18446744073709551615}
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = true
    v8 := []int{510, -2147483648, 4294967295}
    var v9 int8 = 1
    _ = p0
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 uint16, p1 int, p2 uint32) uint8 {
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int = fi11
            var v13 float32 = 3.14159
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := []int{p1, (p1 ^ 2147483647)}
        _ = fi10
        _ = v14
    }
    var v15 uint16 = p0
    v16 := []int{(p1 * 9223372036854775807)}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 239
}

func main() {
    var v1 uint32 = 1
    var v2 bool = true
    var v3 uint32 = v1
    var v4 int8 = 57
    v17 := []int{32768, -780, -399}
    var v18 string = "foo bar"
    var v19 uint = 72
    _ = len(v17)
    if (!(!v2)) {
        var v20 float64 = float64(1.5)
        if v2 {
            v3 = (v1 - (v1 - v1))
        } else {
            var v21 int16 = 0
            var v22 int64 = 158
            fmt.Println(v18)
            _ = v21
            _ = v22
        }
        _ = v20
    } else {
        fmt.Println(v18, v4)
        var v23 uint8 = 1
        _ = v23
    }
    _ = len(v17)
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS1{F0: 1, F1: 42}
        var v26 uint64 = 2
        if v2 {
            if len(v17) > 0 {
                v17[0] = fi24
            }
            v18 = "hello"
            var v27 int32 = 0
            _ = v27
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    v4 += (v4 / ((v4) | 1))
    if false {
        v18 = v18
        var v28 int16 = 42
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS1{F0: genFn0((v28 >= (v28 % ((-1) | 1)))), F1: 0}
            var v31 float32 = 0.1
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    } else {
        var v32 string = v18
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 uint16 = 42
        _ = fi33
        _ = v34
    }
    v4 = v4
    v19 = v19
    var v35 int32 = -1
    fmt.Println(v18, v3)
    if (!(2 > 121)) {
        var v36 int16 = 1000
        for fi37 := 0; fi37 < 8; fi37++ {
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            if true {
                v1 -= (^(v3))
            } else {
                v17 = append(v17, (fi37 / (((fi37 >> (((v19 * v19)) & 63))) | 1)))
                v3 += 42
            }
            for fi38 := 0; fi38 < 1; fi38++ {
                if true {
                    v19 = (v19 / (((v19 + 18446744073709551615)) | 1))
                } else {
                    var v39 int = (fi37 / (((^(fi38))) | 1))
                    v19 = (v19 / (((v19 & v19)) | 1))
                    _ = v39
                }
                if v2 {
                    var v40 bool = ("a" <= "foo bar")
                    var v41 int = (fi38 + 2147483647)
                    _ = v40
                    _ = v41
                }
                var v42 uint16 = 32767
                var v43 uint8 = genFn1((v42 ^ v42), fi37, v3)
                _ = fi38
                _ = v42
                _ = v43
            }
            _ = fi37
        }
        _ = v36
    } else {
        var v44 uint16 = 65535
        var v45 int16 = 42
        var v46 uint16 = (v44 | (v44 / (((v44 | 862)) | 1)))
        _ = v44
        _ = v45
        _ = v46
    }
    fmt.Println(v19, v35, v18)
    v47 := GenS1{F0: 32767, F1: 111}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v47
}

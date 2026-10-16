package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int16 }
func genFn0() int16 {
    var v4 int32 = 2147483647
    if (1 < 18446744073709551615) {
        for fi5 := 0; fi5 < 18; fi5++ {
            fmt.Println(v4, fi5)
            _ = fi5
        }
    } else {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := []int{(fi6 + (fi6 ^ -1)), fi6, (-(fi6))}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi6
            _ = v7
        }
    }
    _ = v4
    return -111
}
func genFn1(p0 uint16) uint8 {
    p0 += (p0 + p0)
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(p0, fi8)
        fmt.Println(fi8, fi8)
        if false {
            fmt.Println(p0, p0)
            for fi9 := 0; fi9 < 3; fi9++ {
                p0 = p0
                _ = fi9
            }
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            p0 += (^(p0))
            if true {
                fmt.Println(fi11, fi11)
                if false {
                    if true {
                        p0 = p0
                        var v12 int32 = -2147483648
                        _ = v12
                    }
                    v13 := []int{}
                    var v14 string = "\n\t"
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(fi10, fi11, p0)
                    var v15 uint64 = 4294967295
                    v16 := GenS1{F0: 127, F1: 754, F2: -601}
                    _ = v15
                    _ = v16
                }
            } else {
                p0 = p0
                var v17 uint64 = 205
                _ = v17
            }
            v18 := []int{(-(fi11))}
            _ = fi11
            _ = v18
        }
        if (true && false) {
            var v19 uint32 = 0
            var v20 string = "hello"
            for fi21 := 0; fi21 < 9; fi21++ {
                v22 := GenS1{F0: 61, F1: 2147483647, F2: 1}
                p0 = 32767
                v23 := []int{(fi10 % ((fi10) | 1)), (fi10 + 42)}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v19
            _ = v20
        } else {
            v24 := []int{-2147483648, (^(fi10))}
            fmt.Println(fi10, p0)
            _ = v24
        }
        _ = fi10
    }
    if false {
        var v25 uint = 646
        var v26 int16 = -76
        v27 := []int{2147483647, 32767, 32768, 32768}
        _ = v25
        _ = v26
        _ = v27
    } else {
        if (!false) {
            var v28 int16 = 0
            var v29 float64 = 0.1
            _ = v28
            _ = v29
        }
        if false {
            var v30 uint8 = 128
            for fi31 := 0; fi31 < 8; fi31++ {
                p0 = (p0 / ((p0) | 1))
                for fi32 := 0; fi32 < 19; fi32++ {
                    fmt.Println(p0)
                    var v33 int = (^(fi32))
                    v34 := []int{(fi32 * (^(fi32))), fi32}
                    var v35 int8 = 100
                    _ = fi32
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = fi31
            }
            _ = v30
        } else {
            var v36 int32 = -830
            _ = v36
        }
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.0
    var v3 bool = ((42 >= 127) || false)
    _ = len(v1)
    _ = len(v1)
    switch "" {
    case "a":
        v37 := []int{-126, -2147483648, 9223372036854775807, -1}
        _ = v37
    case "\n\t":
        v38 := GenS1{F0: 0, F1: 1, F2: genFn0()}
        _ = v38
    }
    v2 -= v2
    for fi39 := 0; fi39 < 1; fi39++ {
        for fi40 := 0; fi40 < 3; fi40++ {
            v41 := GenS1{F0: -76, F1: -1, F2: 32767}
            _ = fi40
            _ = v41
        }
        v2 -= -2.25
        _ = fi39
    }
    for fi42 := 0; fi42 < 1; fi42++ {
        v1["hello"] = "a"
        if (v3 && (275 == 2147483647)) {
            v43 := GenS1{F0: 0, F1: -2147483648, F2: 1000}
            v44 := v1[("foo bar" + "hello")]
            if v3 {
                var v45 int64 = 99
                _ = v45
            }
            _ = v43
            _ = v44
        }
        _ = fi42
    }
    for fi46 := 0; fi46 < 20; fi46++ {
        v2 -= v2
        for fi47 := 0; fi47 < 16; fi47++ {
            var v48 uint64 = 136
            if (false || (1 >= 626)) {
                fmt.Println(fi46)
            } else {
                var v49 bool = v3
                var v50 int32 = 1
                var v51 string = "foo bar"
                _ = v49
                _ = v50
                _ = v51
            }
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    var v52 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
func genFn0(p0 int8, p1 uint, p2 uint32) uint {
    var v5 string = ("hello" + "\n\t")
    v6 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (^(p1))
}

func main() {
    var v1 uint8 = 166
    v2 := []int{}
    var v3 bool = false
    v4 := []string{((((("foo bar" + ((("a" + "") + "\n\t") + "a")) + "a") + ("\n\t" + "")) + "a") + "a"), "\n\t"}
    v1 = (v1 * (v1 * v1))
    var v7 float64 = 1.5
    var v8 float64 = 1e10
    for fi9 := 0; fi9 < 3; fi9++ {
        v3 = v3
        _ = fi9
    }
    switch 652 {
    case -1:
        var v10 int8 = -45
        _ = v10
    }
    if (-1 <= 0) {
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 int64 = 42
            if v3 {
                var v13 int64 = v12
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                fmt.Println(v7, v1, v7)
                _ = v13
            }
            if v3 {
                var v14 uint8 = v1
                var v15 int = (fi11 ^ fi11)
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 float64 = 1.5
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if v3 {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        v14 += 1
                        v2 = append(v2, (v15 + fi16))
                    }
                    _ = fi16
                    _ = v17
                }
                _ = v14
                _ = v15
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 uint64 = 4294967295
                    v7 -= v8
                    var v20 uint16 = 246
                    var v21 float32 = (-2.25 + 0.0)
                    if v3 {
                        switch fi11 {
                        case 689:
                            v1 = 99
                        case 1:
                            fmt.Println(v19, fi11)
                        case -642:
                            v22 := GenS1{F0: 2147483648}
                            _ = v22
                        default:
                            v23 := GenS2{F0: -625}
                            _ = v23
                        }
                    }
                    _ = len(v4)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                v3 = true
            }
            _ = fi11
            _ = v12
        }
    } else {
        var v24 float64 = 0.1
        _ = v24
    }
    var v25 int = -682
    for fi26 := 0; fi26 < 16; fi26++ {
        _ = len(v4)
        v1 = (v1 - (v1 * (v1 + 127)))
        for fi27 := 0; fi27 < 14; fi27++ {
            v7 += 3.14159
            v28 := GenS3{F0: -2147483648}
            v29 := GenS3{F0: 2147483647}
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    if true {
        fmt.Println(v3)
        v30 := GenS4{F0: 70}
        _ = v30
    } else {
        if false {
            var v31 int = v25
            if len(v2) > 0 {
                v2[0] = 127
            }
            _ = v31
        } else {
            fmt.Println(v3, v8)
            if (-1 != 45) {
                v8 += v8
            }
        }
        fmt.Println(v3)
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        v8 += float64(-2.25)
        _ = fi32
    }
    v2 = append(v2, v25)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 int64 = -9223372036854775808
        if v3 {
            v35 := GenS5{F0: 1}
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        if v3 {
            switch "hello" {
            case "":
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            case "hello":
                v25 = v25
            case "\n\t":
                if (!((0 >= 0) && (!v3))) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v37 string = "hello"
                    switch (fi36 * fi36) {
                    case 127:
                        switch (v25 - 4294967295) {
                        case 0:
                            v38 := GenS3{F0: 0}
                            _ = v38
                        case 951:
                            fmt.Println(fi36)
                        default:
                            v8 += (v7 * v8)
                        }
                    case -222:
                        for fi39 := 0; fi39 < 3; fi39++ {
                            var v40 string = "hello"
                            _ = fi39
                            _ = v40
                        }
                    case 32768:
                        fmt.Println(v7, fi36, v7)
                    }
                    _ = v37
                }
            }
        }
        _ = fi36
    }
    var v41 uint8 = 136
    fmt.Println(v1, v3, v3)
    v4 = append(v4, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v7
    _ = v8
}

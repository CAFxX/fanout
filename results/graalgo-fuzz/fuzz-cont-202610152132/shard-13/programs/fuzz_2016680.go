package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 string; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 int64) int16 {
    p0 = 9223372036854775807
    fmt.Println(p0)
    fmt.Println(p0)
    if false {
        v5 := GenS1{F0: 2, F1: 816}
        p0 += (p0 >> ((255) & 63))
        _ = v5
    } else {
        v6 := GenS1{F0: 255, F1: -620}
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 bool = (!true)
            if v8 {
                var v9 float64 = 0.0
                _ = v9
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                if v8 {
                    v11 := GenS2{F0: fi7}
                    var v12 uint16 = 42
                    _ = v11
                    _ = v12
                }
                if (!(-32768 > 0)) {
                    if v8 {
                        var v13 float32 = ((3.14159 / ((1e10 - 0.1) + (1e10 + (0.0 * (((0.1 * -2.25) / -2.25) + 3.14159))))) - float32(float64((-2.25 / 3.14159))))
                        v14 := []int{1}
                        var v15 uint16 = 2
                        fmt.Println(v15)
                        v8 = (v8 && (0.1 > 1e10))
                        _ = len(v14)
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS3{F0: 2, F1: "\n\t", F2: 805}
                    _ = v16
                } else {
                    v17 := GenS4{F0: 127}
                    v18 := []int{fi7, -9223372036854775808, fi7, fi10}
                    if v8 {
                        var v19 uint32 = 65535
                        v18 = append(v18, fi7)
                        v19 = (^(v19))
                        var v20 uint32 = (v19 >> ((94) & 63))
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int16 = -297
                        switch "a" {
                        case "\n\t":
                            var v22 uint32 = 2147483647
                            _ = v22
                        case "foo bar":
                            v8 = true
                        default:
                            var v23 int8 = 100
                            _ = v23
                        }
                        var v24 uint16 = 417
                        _ = v21
                        _ = v24
                    }
                    fmt.Println(fi7)
                    if ((true && ((!(!v8)) || true)) || (!(!v8))) {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    } else {
                        if true {
                            v8 = (0 >= 0)
                            var v25 uint16 = 65535
                            v18 = append(v18, (fi7 >> ((18446744073709551615) & 63)))
                            _ = v25
                        } else {
                            fmt.Println(fi10)
                            var v26 bool = (!v8)
                            v27 := GenS1{F0: 128, F1: 1}
                            _ = v26
                            _ = v27
                        }
                        var v28 string = "hello"
                        _ = v28
                    }
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = v17
                    _ = v18
                }
                v29 := []int{(fi10 << ((1) & 63))}
                v30 := GenS5{F0: 32767}
                var v31 int16 = 1
                _ = fi10
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi7
            _ = v8
        }
        if true {
            v32 := []int{-276, 128, 32768}
            _ = v32
        } else {
            p0 -= 42
            var v33 uint64 = 2
            v34 := []int{32768, -9223372036854775808, -2147483648, 1}
            _ = v33
            _ = v34
        }
        _ = v6
    }
    _ = p0
    return 968
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    var v3 uint = 1
    var v4 float64 = -2.25
    var v35 string = "hello"
    v3 = (v3 << ((v3) & 63))
    v36 := GenS4{F0: v2}
    if true {
        switch v35 {
        case "":
            var v37 bool = (((-2.25 - (1e10 - 0.0)) - 0.0) < 0.1)
            _ = v37
        }
    } else {
        if true {
            v2 += v2
        } else {
            for fi38 := 0; fi38 < 4; fi38++ {
                v35 = "foo bar"
                v3 = v3
                fmt.Println(v2, fi38)
                _ = fi38
            }
            v2 += v2
        }
        var v39 uint64 = 4294967295
        var v40 int64 = 0
        _ = v39
        _ = v40
    }
    var v41 string = ("" + (("" + "foo bar") + "a"))
    fmt.Println(v2, v2)
    v35 = (v41 + (v41 + v35))
    if true {
        for fi42 := 0; fi42 < 9; fi42++ {
            v43 := GenS6{F0: 134}
            v44 := GenS5{F0: 1}
            for fi45 := 0; fi45 < 8; fi45++ {
                var v46 int = fi45
                _ = fi45
                _ = v46
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        var v47 uint16 = 2
        _ = v47
    } else {
        v48 := GenS6{F0: 409}
        _ = v48
    }
    if false {
        switch ((v41 + "\n\t") + "") {
        case "":
            for fi49 := 0; fi49 < 5; fi49++ {
                v50 := v1[((v35 + v41) + "hello")]
                for fi51 := 0; fi51 < 19; fi51++ {
                    v4 -= 1.5
                    _ = len(v1)
                    v1[v35] = int(v2)
                    _ = fi51
                }
                _ = fi49
                _ = v50
            }
        case "hello":
            if false {
                var v52 int8 = -119
                _ = v52
            }
        case "a":
            v35 = ("hello" + v35)
        }
        fmt.Println(v3, v2)
        v4 = v4
    }
    var v53 float64 = 1e10
    var v54 int = -1
    fmt.Println(v54, v41)
    for fi55 := 0; fi55 < 19; fi55++ {
        v53 -= v4
        var v56 uint8 = 180
        _ = fi55
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v53
    _ = v54
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 int8; F1 int16 }

func main() {
    var v1 uint8 = 128
    var v2 int32 = -808
    var v3 bool = true
    var v4 uint16 = 255
    var v5 int = 385
    var v6 int64 = 42
    for fi7 := 0; fi7 < 18; fi7++ {
        if v3 {
            var v8 uint = 65535
            _ = v8
        } else {
            var v9 uint16 = v4
            switch ("\n\t" + "") {
            case "hello":
                var v10 int64 = 9223372036854775807
                _ = v10
            default:
                v2 = (v2 * 1)
            }
            v4 -= 42
            _ = v9
        }
        v3 = v3
        v11 := []int{fi7}
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: -1, F1: 1, F2: (v1 % ((v1) | 1))}
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int16 = 165
            if ((v2 / ((v2) | 1)) >= v2) {
                if (!true) {
                    fmt.Println(v15)
                    v16 := []int{}
                    var v17 uint = 42
                    _ = v16
                    _ = v17
                } else {
                    if false {
                        v6 -= (v6 + 246)
                    }
                    if (v4 <= 42) {
                        v15 = (v15 % (((v15 + int16(v4))) | 1))
                    }
                }
                var v18 uint16 = v4
                _ = v18
            }
            if false {
                if v3 {
                    v1 = (v1 % (((v1 ^ v1)) | 1))
                    var v19 bool = true
                    var v20 string = "\n\t"
                    _ = v19
                    _ = v20
                } else {
                    switch fi12 {
                    case 42:
                        v21 := []int{}
                        v6 = v6
                        _ = v21
                    }
                    var v22 string = "\n\t"
                    _ = v22
                }
            } else {
                v6 = v6
                v15 += v15
            }
            var v23 uint64 = 846
            var v24 uint64 = v23
            var v25 string = "foo bar"
            _ = fi14
            _ = v15
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi12
        _ = v13
    }
    var v26 uint32 = 0
    var v27 bool = (!true)
    v28 := GenS2{F0: (-(v2))}
    var v29 bool = (255 <= v4)
    var v30 float32 = (3.14159 + 1.5)
    if (!v29) {
        switch v5 {
        case -997:
            v1 += 1
        case 1:
            v31 := []int{(v5 >> ((uint(v1)) & 63)), v5, v5, (v5 | v5)}
            _ = v31
        case 955:
            var v32 int16 = 42
            _ = v32
        default:
            v26 += (v26 + (v26 >> ((18446744073709551615) & 63)))
        }
        v33 := GenS2{F0: 1}
        v30 -= v30
        _ = v33
    }
    fmt.Println(v30, v29, v1)
    v34 := GenS3{F0: (v26 - v26), F1: v4}
    switch "\n\t" {
    case "":
        v35 := GenS4{F0: 1, F1: 32767}
        _ = v35
    case "hello":
        v36 := GenS2{F0: 1}
        _ = v36
    }
    fmt.Println(v2, v5, v5)
    v37 := []int{v5, int(v4)}
    for fi38 := 0; fi38 < 18; fi38++ {
        if v27 {
            for fi39 := 0; fi39 < 4; fi39++ {
                if (!v3) {
                    var v40 uint64 = 2
                    var v41 int32 = (v2 * v2)
                    _ = v40
                    _ = v41
                }
                var v42 float32 = -2.25
                _ = fi39
                _ = v42
            }
            v37 = append(v37, 0)
            for fi43 := 0; fi43 < 2; fi43++ {
                var v44 uint32 = 969
                var v45 float64 = 1.5
                switch v5 {
                case 4294967295:
                    v37 = append(v37, 2147483647)
                default:
                    if v3 {
                        var v46 uint16 = (^(v4))
                        _ = v46
                    } else {
                        switch (fi43 << ((948) & 63)) {
                        case 1:
                            fmt.Println(v4)
                        case -1:
                            fmt.Println(fi43)
                        case 42:
                            v26 -= v26
                        default:
                            var v47 float64 = (3.14159 + v45)
                            _ = v47
                        }
                        fmt.Println(fi38, v6, v6)
                    }
                }
                v48 := GenS2{F0: v2}
                v27 = (-471 < 42)
                _ = fi43
                _ = v44
                _ = v45
                _ = v48
            }
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint16 }
func genFn0() int16 {
    var v5 int64 = 9223372036854775807
    var v6 int32 = -575
    _ = v5
    _ = v6
    return 32767
}
func genFn1(p0 int16) int8 {
    v7 := []int{2147483647, int(p0)}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 1
    var v2 uint = 1
    var v3 bool = true
    var v4 uint32 = 1
    fmt.Println(v4)
    v8 := GenS1{F0: -52}
    fmt.Println(v2)
    var v9 uint = 797
    if (-212 > -315) {
        v10 := GenS1{F0: 2147483647}
        v4 = 2147483647
        if v3 {
            fmt.Println(v9, v9)
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS2{F0: int64(v4)}
                if (!(255 < 42)) {
                    var v13 int16 = int16(v9)
                    var v14 bool = false
                    _ = v13
                    _ = v14
                }
                switch "hello" {
                case "hello":
                    v15 := []int{(fi11 >> (((v2 >> ((v9) & 63))) & 63)), (fi11 << ((1) & 63)), (^(fi11)), (fi11 % (((fi11 << (((v2 >> ((v2) & 63))) & 63))) | 1))}
                    _ = v15
                case "\n\t":
                    if v3 {
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v17 := GenS1{F0: 1}
                            _ = fi16
                            _ = v17
                        }
                        var v18 int64 = 543
                        v19 := []int{-539, fi11, fi11}
                        _ = v18
                        _ = v19
                    }
                }
                _ = fi11
                _ = v12
            }
        }
        _ = v10
    }
    var v20 string = ("a" + "a")
    switch int(v4) {
    case 90:
        switch v20 {
        case "hello":
            v21 := GenS3{F0: v4}
            _ = v21
        case "\n\t":
            for fi22 := 0; fi22 < 12; fi22++ {
                if v3 {
                    v23 := []int{(fi22 & (fi22 << (((v2 - 390)) & 63))), (fi22 / (((fi22 | fi22)) | 1))}
                    _ = v23
                }
                v24 := []int{(fi22 * (fi22 + fi22))}
                v4 += (v4 + v4)
                _ = fi22
                _ = v24
            }
        case "a":
            var v25 uint8 = v1
            _ = v25
        default:
            v9 = v2
        }
    case 9223372036854775807:
        v26 := []int{2147483647, 179, 32768}
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        if false {
            for fi28 := 0; fi28 < 7; fi28++ {
                switch "a" {
                case "":
                    var v29 uint64 = 201
                    _ = v29
                }
                if true {
                    switch (v20 + ("a" + "hello")) {
                    case "hello":
                        v30 := []int{(fi27 % ((fi27) | 1))}
                        _ = v30
                    default:
                        switch fi27 {
                        case 171:
                            v4 = v4
                        default:
                            fmt.Println(v20, v9, v1)
                        }
                    }
                } else {
                    v20 = v20
                }
                fmt.Println(fi28, v4, v3)
                _ = fi28
            }
        }
        v20 = "foo bar"
        var v31 int8 = 13
        _ = fi27
        _ = v31
    }
    var v32 uint32 = v4
    if false {
        if v3 {
            v1 = 0
        } else {
            var v33 int = 861
            for fi34 := 0; fi34 < 19; fi34++ {
                v35 := []int{fi34, (v33 + (^(fi34))), fi34, -101}
                _ = fi34
                _ = v35
            }
            _ = v33
        }
        fmt.Println(v20, v2)
        var v36 int32 = -1
        _ = v36
    } else {
        for fi37 := 0; fi37 < 12; fi37++ {
            fmt.Println(v20)
            for fi38 := 0; fi38 < 5; fi38++ {
                if v3 {
                    switch "" {
                    case "hello":
                        v2 -= 65535
                    }
                }
                v3 = (!((v32 * v4) > v32))
                if (v3 && v3) {
                    v39 := []int{-2147483648, 571, -283, (fi38 - (-(fi38)))}
                    var v40 float32 = ((-2.25 / 1.5) + float32((1e10 * 1e10)))
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                    _ = v39
                    _ = v40
                } else {
                    var v41 int16 = -1
                    _ = v41
                }
                v42 := []int{int(v9), 127}
                _ = fi38
                _ = v42
            }
            v3 = false
            _ = fi37
        }
        fmt.Println(v20, v1)
        switch ((v20 + (v20 + ("\n\t" + ((v20 + v20) + ("a" + v20))))) + "\n\t") {
        case "a":
            var v43 string = "hello"
            _ = v43
        default:
            if ((v20 + "foo bar") >= (v20 + v20)) {
                v44 := GenS4{F0: v9, F1: (v20 + v20)}
                _ = v44
            }
        }
    }
    switch v20 {
    case "a":
        var v45 int = -265
        _ = v45
    case "":
        var v46 string = ""
        _ = v46
    case "hello":
        v20 = "hello"
    default:
        if (!v3) {
            v47 := GenS5{F0: 610}
            _ = v47
        } else {
            switch v20 {
            case "\n\t":
                v32 -= v32
            case "foo bar":
                for fi48 := 0; fi48 < 4; fi48++ {
                    v3 = (v3 && ((float32((0.1 + (0.0 + 1.5))) != 3.14159) && (!(255 <= 1))))
                    v1 = v1
                    v49 := GenS1{F0: 1}
                    _ = fi48
                    _ = v49
                }
            default:
                v50 := []int{-108, 113}
                _ = v50
            }
            var v51 uint8 = v1
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v8
    _ = v9
}

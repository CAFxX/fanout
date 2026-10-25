package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 string; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 string }
func genFn0(p0 int16) int16 {
    p0 -= -1
    v7 := []int{}
    _ = len(v7)
    _ = p0
    _ = v7
    return p0
}

func main() {
    var v1 bool = (true || false)
    var v2 int8 = 31
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    v5 := []uint32{}
    var v6 uint64 = 734
    if v1 {
        var v8 int64 = 171
        v2 += 100
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: "a", F1: fi9}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        _ = v8
    }
    switch ("" + "hello") {
    case "":
        v2 -= (v2 & 127)
    case "hello":
        for fi12 := 0; fi12 < 2; fi12++ {
            switch 803 {
            case 42:
                if v1 {
                    var v13 uint8 = 128
                    var v14 int16 = -32768
                    _ = v13
                    _ = v14
                } else {
                    v1 = v1
                    switch "foo bar" {
                    case "\n\t":
                        switch "\n\t" {
                        case "foo bar":
                            v6 += (v6 - v6)
                        case "hello":
                            var v15 uint16 = 0
                            _ = v15
                        }
                    case "a":
                        var v16 int8 = v2
                        _ = v16
                    }
                    v17 := GenS2{F0: "\n\t", F1: 2, F2: 113}
                    _ = v17
                }
            case 71:
                switch fi12 {
                case 0:
                    v18 := GenS2{F0: ("a" + "foo bar"), F1: 2147483648, F2: 40}
                    _ = v18
                case 824:
                    var v19 uint16 = 255
                    _ = v19
                default:
                    var v20 bool = v1
                    _ = v20
                }
            }
            if (float32((1.5 / 1.0)) >= float32(0.0)) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    _ = len(v3)
                    var v22 uint16 = 0
                    _ = fi21
                    _ = v22
                }
                var v23 uint = 255
                _ = v23
            } else {
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint = 4294967295
                    if (255 < 127) {
                        fmt.Println(v1, fi12)
                        switch (("a" + "foo bar") + ("\n\t" + ("\n\t" + ""))) {
                        case "foo bar":
                            fmt.Println(v1)
                        case "\n\t":
                            if v1 {
                                v2 = v2
                            }
                        case "":
                            if (false || v1) {
                                fmt.Println(fi12)
                            } else {
                                v26 := GenS3{F0: 42, F1: 1, F2: ("hello" + "hello")}
                                delete(v3, 260)
                                v1 = ((!(!(!v1))) && false)
                                _ = v26
                            }
                        }
                        var v27 int = (fi24 + -313)
                        var v28 int16 = 1000
                        _ = v27
                        _ = v28
                    }
                    var v29 bool = ((v6 % (((v6 & v6)) | 1)) > (^(v6)))
                    _ = fi24
                    _ = v25
                    _ = v29
                }
                v30 := v3[2]
                var v31 int32 = -19
                _ = v30
                _ = v31
            }
            _ = fi12
        }
    }
    v32 := GenS3{F0: 0, F1: -177, F2: "foo bar"}
    var v33 uint32 = 42
    for fi34 := 0; fi34 < 6; fi34++ {
        v33 = 4294967295
        _ = fi34
    }
    var v35 uint32 = (v33 / (((v33 - v33)) | 1))
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := GenS3{F0: int16(v2), F1: -2147483648, F2: "a"}
            _ = fi37
            _ = v38
        }
        switch fi36 {
        case 127:
            if (v1 || ((!(255 != 899)) || (2 >= 18446744073709551615))) {
                var v39 uint64 = (v6 % (((v6 & v6)) | 1))
                var v40 int32 = 955
                var v41 float32 = float32(3.14159)
                _ = v39
                _ = v40
                _ = v41
            } else {
                if v1 {
                    if ((int16(v33) > genFn0(int16(fi36))) && (!v1)) {
                        if (v1 && (0.0 == 1e10)) {
                            var v42 bool = (1e10 <= (3.14159 - -2.25))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            for fi43 := 0; fi43 < 4; fi43++ {
                                v6 = v6
                                _ = fi43
                            }
                            _ = v42
                        }
                        var v44 int16 = 32767
                        if len(v4) > 0 {
                            v4[0] = ("foo bar" + "foo bar")
                        }
                        _ = v44
                    }
                    switch ("hello" + "hello") {
                    case "":
                        for fi45 := 0; fi45 < 8; fi45++ {
                            v1 = v1
                            v5 = append(v5, 2)
                            _ = fi45
                        }
                    case "hello":
                        delete(v3, (v35 & (^(v35))))
                    }
                }
                var v46 string = "a"
                _ = v46
            }
        }
        _ = fi36
    }
    delete(v3, v35)
    switch "\n\t" {
    case "\n\t":
        v3[v33] = (v33 + 4294967295)
    }
    var v47 uint16 = 32767
    var v48 float64 = float64(((-2.25 * -2.25) / 1.0))
    fmt.Println(v35, v33, v33)
    fmt.Println(v48, v48)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}

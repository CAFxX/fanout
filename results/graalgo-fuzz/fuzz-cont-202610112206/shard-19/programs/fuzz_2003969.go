package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint8 = 152
    v2 := make(map[int]bool)
    var v3 int = 1
    var v4 int = (v3 / (((v3 + v3)) | 1))
    var v5 float32 = 0.0
    if (3.14159 != -2.25) {
        _ = len(v2)
        v5 = v5
        var v6 int64 = 1
        _ = v6
    }
    v2[(v4 * v3)] = false
    fmt.Println(v4, v4, v1)
    v7 := GenS1{F0: (1 <= -1)}
    var v8 int8 = -103
    var v9 int64 = 42
    v5 -= v5
    v9 = (v9 << ((65535) & 63))
    v1 -= v1
    if ((!(117 <= (v1 | 42))) && ((-2.25 / 1e10) < float64(0.0))) {
        fmt.Println(v5)
    } else {
        var v10 int8 = v8
        for fi11 := 0; fi11 < 20; fi11++ {
            if false {
                var v12 bool = (("a" > "") || (!true))
                fmt.Println(v12, v3)
                fmt.Println(v9, v9)
                _ = v12
            }
            fmt.Println(v5, v9)
            var v13 string = "\n\t"
            _ = fi11
            _ = v13
        }
        var v14 int8 = v10
        _ = v10
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := []int{(v3 | (fi15 & 42))}
        if false {
            if (("\n\t" >= ("a" + "")) && true) {
                var v17 int64 = (v9 / ((42) | 1))
                fmt.Println(fi15)
                if (198 < 42) {
                    v3 += v3
                    var v18 uint = 255
                    _ = v18
                }
                _ = v17
            }
            if false {
                fmt.Println(v1)
                switch (v4 >> ((42) & 63)) {
                case 127:
                    switch 32767 {
                    case -878:
                        v5 = v5
                    default:
                        v5 -= 1e10
                    }
                default:
                    v2[fi15] = false
                }
            } else {
                var v19 uint16 = 65535
                for fi20 := 0; fi20 < 1; fi20++ {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    _ = fi20
                }
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 uint32 = 520
                    _ = fi21
                    _ = v22
                }
                _ = v19
            }
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 uint8 = (^(v1))
            switch 32767 {
            case 4294967295:
                if true {
                    if ((false || false) && (!true)) {
                        v25 := GenS1{F0: ((-(v3)) < v3)}
                        switch "\n\t" {
                        case "a":
                            v5 += (v5 * (((((float32((1.5 - 0.1)) + v5) * float32(0.1)) + -2.25) * (v5 + v5)) * float32(((3.14159 + 0.1) + -2.25))))
                        case "\n\t":
                            if true {
                                v24 += (v1 / (((v24 >> ((2) & 63))) | 1))
                                v26 := GenS1{F0: (3.14159 == 1.5)}
                                var v27 uint16 = 32767
                                _ = v26
                                _ = v27
                            } else {
                                var v28 float32 = ((float32(0.0) - 0.0) - -2.25)
                                _ = v28
                            }
                        default:
                            fmt.Println(v5, fi15)
                        }
                        v5 += v5
                        _ = v25
                    } else {
                        if (42 != 72) {
                            switch (v4 >> ((42) & 63)) {
                            case 32767:
                                switch "foo bar" {
                                case "":
                                    v16 = append(v16, 161)
                                }
                            case 32768:
                                var v29 uint32 = 2
                                _ = v29
                            }
                            if (-2.25 >= ((float64(3.14159) / float64((v5 + 1e10))) * 1e10)) {
                                if (-32768 != 1000) {
                                    v24 = (v1 | 2)
                                }
                            }
                            switch (^(fi15)) {
                            case 128:
                                fmt.Println(fi23, fi23, fi23)
                                var v30 int = 2147483647
                                v31 := GenS1{F0: false}
                                _ = v30
                                _ = v31
                            default:
                                var v32 int = 4294967295
                                _ = v32
                            }
                        } else {
                            var v33 uint32 = 42
                            delete(v2, fi15)
                            v34 := GenS2{F0: (!true)}
                            _ = v33
                            _ = v34
                        }
                        v35 := v2[9223372036854775807]
                        v36 := GenS2{F0: false}
                        _ = v35
                        _ = v36
                    }
                    v1 -= v1
                    v37 := GenS2{F0: true}
                    if (!(-2147483648 != 60)) {
                        v5 = (float32((((3.14159 - (1.5 - 0.1)) + (float64(v5) + -2.25)) + 3.14159)) / (v5 / (0.0 + (v5 + v5))))
                    }
                    v9 -= v9
                    _ = v37
                }
            case 128:
                switch 9223372036854775807 {
                case 8:
                    v38 := GenS3{F0: 176}
                    _ = v38
                case -808:
                    v5 = v5
                case -748:
                    if len(v16) > 0 {
                        v16[0] = (fi15 << ((0) & 63))
                    }
                }
            }
            _ = fi23
            _ = v24
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}

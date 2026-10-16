package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint64; F2 uint32 }
func genFn0(p0 uint64, p1 uint16, p2 string) uint8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 string = p2
        v7 := GenS1{F0: 1}
        var v8 bool = (0 >= 0)
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        switch p2 {
        case "":
            fmt.Println(p2, p2)
        case "a":
            var v10 int64 = 42
            _ = v10
        default:
            v11 := GenS1{F0: -32768}
            _ = v11
        }
        switch (p2 + (p2 + "")) {
        case "foo bar":
            for fi12 := 0; fi12 < 5; fi12++ {
                p1 -= (^(p1))
                _ = fi12
            }
        case "\n\t":
            if (!(!false)) {
                var v13 int16 = 10
                p2 = p2
                _ = v13
            } else {
                v14 := GenS1{F0: -818}
                _ = v14
            }
        case "a":
            v15 := []int{}
            _ = v15
        }
        p1 = uint16(p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    v1 := []int{127, 9223372036854775807, 32767}
    var v2 int64 = 0
    v3 := []string{"foo bar", ("foo bar" + ("hello" + ("" + ("hello" + ("foo bar" + "foo bar")))))}
    v4 := make(map[uint32]uint32)
    v2 -= v2
    var v16 float64 = (1.5 - -2.25)
    if (0.1 <= (v16 - v16)) {
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(fi17, v16, v16)
            _ = fi17
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            if true {
                v19 := GenS2{F0: 42, F1: 780}
                v16 -= v16
                _ = v19
            }
            v20 := v4[948]
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 string = "hello"
                var v23 int32 = -240
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi18
            _ = v20
        }
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        switch "\n\t" {
        case "hello":
            if false {
                switch fi24 {
                case 0:
                    v16 -= (v16 * v16)
                case 32767:
                    v2 -= -708
                case 4294967295:
                    v25 := GenS3{F0: 1, F1: 0, F2: 32767}
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 17; fi26++ {
                        var v27 uint16 = 23
                        if ((!(v2 > 975)) && (0.1 < float32(v16))) {
                            var v28 int = fi26
                            v2 = v2
                            if true {
                                var v29 float64 = (v16 - 0.1)
                                fmt.Println(v2, fi24, v28)
                                if ("hello" < "a") {
                                    v29 = (v16 - float64(float32(v29)))
                                    var v30 uint64 = 42
                                    v27 -= v27
                                    _ = v30
                                }
                                _ = v29
                            } else {
                                v27 = v27
                                fmt.Println(v16)
                            }
                            var v31 int = v28
                            _ = v28
                            _ = v31
                        } else {
                            var v32 string = ""
                            _ = v32
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        case "hello":
                            _ = len(v3)
                        case "a":
                            var v33 string = ""
                            _ = v33
                        default:
                            fmt.Println(v27, fi26, v2)
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                var v34 uint16 = 2
                _ = v34
            } else {
                var v35 float64 = v16
                v16 -= (1.5 * v16)
                v36 := GenS4{F0: 42, F1: 1, F2: 42}
                _ = v35
                _ = v36
            }
        case "":
            var v37 uint32 = 2147483648
            _ = v37
        }
        _ = fi24
    }
    var v38 uint16 = 808
    v38 -= (v38 / (((v38 - 42)) | 1))
    v38 += 647
    for fi39 := 0; fi39 < 3; fi39++ {
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 int32 = 2147483647
            _ = fi40
            _ = v41
        }
        v4[0] = 65535
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}

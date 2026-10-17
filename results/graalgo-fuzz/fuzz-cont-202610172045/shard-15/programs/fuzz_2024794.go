package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -1
    var v3 float64 = 1.5
    var v4 int32 = (-(v2))
    if true {
        for fi5 := 0; fi5 < 17; fi5++ {
            if (127 > 4) {
                var v6 int16 = 231
                for fi7 := 0; fi7 < 17; fi7++ {
                    var v8 int32 = v4
                    _ = fi7
                    _ = v8
                }
                v9 := GenS1{F0: "", F1: 65535}
                _ = v6
                _ = v9
            } else {
                for fi10 := 0; fi10 < 15; fi10++ {
                    v3 += 0.1
                    fmt.Println(fi10, fi5, v3)
                    _ = fi10
                }
                for fi11 := 0; fi11 < 1; fi11++ {
                    v12 := []int{(fi5 % (((fi11 * fi5)) | 1))}
                    _ = fi11
                    _ = v12
                }
                switch ("hello" + "hello") {
                case "a":
                    var v13 bool = ((!false) || true)
                    _ = v13
                case "hello":
                    if true {
                        var v14 int64 = -1
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            var v16 string = ""
                            var v17 uint16 = 32767
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        var v18 uint64 = 770
                        var v19 uint16 = 32767
                        _ = v18
                        _ = v19
                    }
                case "\n\t":
                    fmt.Println(v2, v2, fi5)
                default:
                    v20 := GenS2{F0: 42, F1: 4294967295}
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 15; fi21++ {
                v3 -= v3
                _ = fi21
            }
            _ = fi5
        }
    } else {
        v4 += (-(v4))
        fmt.Println(v2)
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(fi23)
            v24 := GenS2{F0: -32768, F1: 531}
            var v25 string = ("a" + "foo bar")
            var v26 int16 = -739
            fmt.Println(v25, v25)
            v27 := GenS3{F0: (v26 % ((-32768) | 1)), F1: (fi23 >> ((42) & 63))}
            v28 := []int{}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
            _ = v28
        }
        v29 := GenS1{F0: "\n\t", F1: 32767}
        for fi30 := 0; fi30 < 1; fi30++ {
            v4 -= v2
            _ = fi30
        }
        _ = fi22
        _ = v29
    }
    var v31 int16 = 1
    v32 := v1["foo bar"]
    v33 := GenS4{F0: -254, F1: 4294967295}
    var v34 uint16 = 65535
    if ((!true) && true) {
        v35 := []int{-76, 0}
        v3 = (v3 * (v3 * 1e10))
        _ = v35
    }
    v36 := []int{522, -483}
    v4 += (v4 | (v4 % (((v2 ^ (v2 << ((255) & 63)))) | 1)))
    v32 = v32
    switch 2147483647 {
    case 0:
        switch 128 {
        case -9223372036854775808:
            var v37 int64 = 1
            _ = v37
        case 32768:
            v4 = (v2 / ((v2) | 1))
        default:
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
        }
    }
    v2 -= -748
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    var v38 uint8 = 1
    v4 = v2
    v32 = (v32 + ((v32 + v32) + v32))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v38
    _ = v4
}

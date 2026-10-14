package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint; F1 bool }

func main() {
    var v1 uint32 = 2
    var v2 float32 = 1e10
    var v3 uint64 = 759
    v4 := []string{}
    v5 := GenS1{F0: 543, F1: -1}
    v3 = uint64(v1)
    if false {
        v2 = (v2 + v2)
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 255, F1: 9223372036854775807}
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 uint8 = 255
                v10 := GenS1{F0: 68, F1: -1}
                v11 := GenS2{F0: 65535}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint64 = (v3 / ((18446744073709551615) | 1))
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: 605, F1: 0}
        for fi16 := 0; fi16 < 11; fi16++ {
            switch (fi16 << ((1) & 63)) {
            case 128:
                var v17 uint16 = 65535
                _ = v17
            case -9223372036854775808:
                var v18 uint64 = (v3 & v3)
                _ = v18
            default:
                v19 := GenS2{F0: 1}
                _ = v19
            }
            if (128 >= uint8(v3)) {
                var v20 int = 127
                v21 := GenS2{F0: 32767}
                if false {
                    switch (v20 % ((int(v1)) | 1)) {
                    case 32768:
                        v4 = append(v4, "foo bar")
                        var v22 uint = 98
                        _ = v22
                    case 1:
                        if (0.0 >= 1.5) {
                            v23 := GenS2{F0: 65535}
                            var v24 string = "foo bar"
                            var v25 int32 = int32(v3)
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    }
                }
                var v26 uint64 = v3
                _ = v20
                _ = v21
                _ = v26
            }
            v3 = (v3 * v3)
            switch "foo bar" {
            case "hello":
                var v27 int = fi16
                _ = v27
            case "a":
                v28 := GenS3{F0: (-(fi16)), F1: (v1 % (((v1 & 8)) | 1)), F2: -32768}
                _ = v28
            case "\n\t":
                switch "foo bar" {
                case "":
                    v4 = append(v4, "foo bar")
                case "a":
                    var v29 bool = false
                    _ = v29
                }
            }
            _ = fi16
        }
        var v30 float64 = 0.1
        _ = fi14
        _ = v15
        _ = v30
    }
    var v31 float64 = 0.1
    if false {
        var v32 uint16 = 65535
        var v33 int32 = 2147483647
        v34 := GenS3{F0: -9223372036854775808, F1: (v1 << ((uint(v1)) & 63)), F2: 42}
        _ = v32
        _ = v33
        _ = v34
    } else {
        var v35 uint32 = v1
        var v36 uint = 18446744073709551615
        _ = v35
        _ = v36
    }
    if false {
        v3 -= (^(v3))
        var v37 uint16 = 582
        var v38 string = "a"
        _ = v37
        _ = v38
    } else {
        v39 := GenS4{F0: 816, F1: (false || (!false))}
        v31 += ((float64(-2.25) / 0.1) - v31)
        _ = v39
    }
    var v40 int32 = -2147483648
    v2 = 0.0
    fmt.Println(v40, v31, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v5
}

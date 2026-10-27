package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32 }

func main() {
    var v1 int8 = 100
    var v2 uint16 = 2
    v3 := make(map[uint32]uint32)
    switch ("" + "foo bar") {
    case "":
        var v4 uint = 988
        _ = v4
    case "\n\t":
        var v5 uint16 = (v2 + (v2 << ((804) & 63)))
        _ = v5
    case "hello":
        var v6 bool = (281 > 505)
        _ = v6
    default:
        var v7 int16 = 0
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = 0
        var v10 int64 = -123
        v1 -= v1
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 -= v2
    _ = len(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        v1 = 42
        v12 := []int{fi11, fi11, 837}
        switch (-(fi11)) {
        case 9223372036854775807:
            v13 := v3[65535]
            _ = v13
        case 42:
            v3[2147483648] = 2
        }
        _ = fi11
        _ = v12
    }
    switch "foo bar" {
    case "hello":
        if false {
            if (true && false) {
                v2 = v2
            } else {
                fmt.Println(v1, v1)
            }
            v14 := GenS1{F0: -32768}
            var v15 string = ""
            _ = v14
            _ = v15
        }
    case "foo bar":
        switch "a" {
        case "":
            switch ("foo bar" + ("" + "")) {
            case "foo bar":
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 uint16 = 255
                    for fi18 := 0; fi18 < 14; fi18++ {
                        if false {
                            v17 -= v17
                            v2 = (v17 * v2)
                        }
                        _ = fi18
                    }
                    v3[2] = 2147483648
                    _ = fi16
                    _ = v17
                }
            case "":
                var v19 uint8 = 0
                _ = v19
            }
        case "a":
            if true {
                if true {
                    delete(v3, 559)
                    var v20 float32 = 1e10
                    if true {
                        var v21 uint16 = 708
                        var v22 uint32 = 2147483648
                        if false {
                            fmt.Println(v2)
                            var v23 string = "\n\t"
                            _ = v23
                        }
                        _ = v21
                        _ = v22
                    }
                    _ = v20
                } else {
                    fmt.Println(v1, v2)
                }
                v1 -= 42
                for fi24 := 0; fi24 < 12; fi24++ {
                    v25 := GenS2{F0: -603, F1: 1}
                    v26 := GenS2{F0: -32768, F1: 0}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 int32 = -187
                var v28 int32 = v27
                _ = v27
                _ = v28
            }
        }
    case "\n\t":
        switch 128 {
        case 698:
            if (false && false) {
                var v29 uint64 = 808
                _ = v29
            } else {
                v2 += (v2 + v2)
                var v30 uint32 = 916
                for fi31 := 0; fi31 < 6; fi31++ {
                    if true {
                        var v32 uint16 = 65535
                        v2 -= 223
                        _ = v32
                    }
                    v30 -= v30
                    _ = fi31
                }
                _ = v30
            }
        case -236:
            v3[1] = 182
        }
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 int64 = 502
        var v35 uint8 = 122
        v36 := []int{fi33}
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}

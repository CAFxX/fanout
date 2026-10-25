package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []uint32{42, 65535, 2147483647}
    var v2 uint64 = 845
    var v3 string = (("foo bar" + ((("" + ("" + "")) + "") + "hello")) + "foo bar")
    var v4 uint16 = 404
    for fi5 := 0; fi5 < 4; fi5++ {
        v4 += 0
        _ = fi5
    }
    if (2 != 65535) {
        v4 = (v4 ^ v4)
    } else {
        var v6 int32 = 1
        _ = v6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 := GenS1{F0: "a"}
    fmt.Println(v4, v2)
    if false {
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                v4 += (v4 >> ((42) & 63))
                _ = fi9
            }
            v3 = ""
            v10 := GenS1{F0: v3}
            _ = fi8
            _ = v10
        }
        var v11 int64 = 9223372036854775807
        _ = v11
    }
    var v12 string = (v3 + "")
    for fi13 := 0; fi13 < 16; fi13++ {
        switch -9223372036854775808 {
        case -390:
            switch "foo bar" {
            case "":
                fmt.Println(v4, v12)
            case "\n\t":
                fmt.Println(v4, v3, v2)
            case "hello":
                v3 = ("a" + (((v3 + v12) + "foo bar") + ("\n\t" + ("foo bar" + v12))))
            }
        case -2147483648:
            var v14 uint8 = 128
            _ = v14
        }
        var v15 uint32 = 1
        _ = fi13
        _ = v15
    }
    v16 := GenS1{F0: v3}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v17 uint32 = 833
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 uint64 = v2
        var v20 float32 = 0.1
        v3 = "foo bar"
        _ = fi18
        _ = v19
        _ = v20
    }
    v12 = v12
    v1 = append(v1, (^(v17)))
    if (-2147483648 < 32767) {
        v21 := GenS2{F0: 0, F1: 2147483647}
        v22 := GenS1{F0: v12}
        if false {
            for fi23 := 0; fi23 < 14; fi23++ {
                v1 = append(v1, v17)
                _ = fi23
            }
            if (42 == 1) {
                for fi24 := 0; fi24 < 1; fi24++ {
                    var v25 uint8 = 1
                    var v26 int8 = 25
                    fmt.Println(v17, fi24, v4)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        }
        _ = v21
        _ = v22
    } else {
        if true {
            v2 = uint64(v4)
            if (v4 != v4) {
                for fi27 := 0; fi27 < 16; fi27++ {
                    if (false && (-1 >= 1)) {
                        fmt.Println(v4)
                    }
                    _ = fi27
                }
                var v28 string = ""
                _ = v28
            } else {
                var v29 uint = 18446744073709551615
                if (true && false) {
                    var v30 float32 = 1.5
                    switch v3 {
                    case "\n\t":
                        v31 := GenS2{F0: 42, F1: 1}
                        _ = v31
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    switch -655 {
                    case 9223372036854775807:
                        v32 := GenS3{F0: v17}
                        _ = v32
                    case 127:
                        fmt.Println(v2, v17, v4)
                    }
                    _ = v30
                }
                _ = v29
            }
            if (128 <= 0) {
                var v33 int32 = -2147483648
                var v34 int16 = -463
                _ = v33
                _ = v34
            } else {
                var v35 uint64 = 18446744073709551615
                v4 += 692
                _ = v35
            }
        } else {
            var v36 uint32 = v17
            v4 += (v4 >> ((797) & 63))
            _ = v36
        }
    }
    fmt.Println(v3)
    switch v3 {
    case "a":
        if (true && (!(true || (false && true)))) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v37 uint32 = (v17 % ((65535) | 1))
            fmt.Println(v37, v12)
            var v38 uint16 = (v4 % ((v4) | 1))
            _ = v37
            _ = v38
        }
    default:
        var v39 int64 = -301
        _ = v39
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}

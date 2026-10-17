package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint8 = 42
    var v2 string = "\n\t"
    v3 := []uint32{69, 2}
    v4 := []int{127, -9223372036854775808}
    v2 = v2
    if true {
        v2 = "foo bar"
        for fi5 := 0; fi5 < 19; fi5++ {
            if false {
                for fi6 := 0; fi6 < 14; fi6++ {
                    if false {
                        if len(v4) > 0 {
                            v4[0] = (fi5 >> ((2) & 63))
                        }
                        v1 += v1
                        switch (fi5 % (((fi6 - 974)) | 1)) {
                        case 1:
                            var v7 int16 = 74
                            _ = v7
                        }
                        v2 = (v2 + (v2 + ("foo bar" + v2)))
                        v1 = (v1 & v1)
                    }
                    _ = fi6
                }
            } else {
                fmt.Println(v2)
                for fi8 := 0; fi8 < 7; fi8++ {
                    v3 = append(v3, 165)
                    v2 = ""
                    var v9 uint16 = 1
                    if false {
                        switch v2 {
                        case "a":
                            v1 = (v1 - (v1 >> ((724) & 63)))
                        case "\n\t":
                            var v10 uint32 = 65535
                            _ = v10
                        }
                        var v11 uint8 = (v1 | (v1 ^ v1))
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
            }
            v3 = append(v3, 65535)
            _ = fi5
        }
    } else {
        switch -647 {
        case -2147483648:
            v12 := GenS1{F0: 18446744073709551615}
            _ = v12
        case -1:
            for fi13 := 0; fi13 < 1; fi13++ {
                fmt.Println(fi13)
                var v14 int = (-(fi13))
                _ = fi13
                _ = v14
            }
        default:
            v2 = v2
        }
        var v15 uint32 = 2147483647
        var v16 string = v2
        _ = v15
        _ = v16
    }
    v1 = (v1 | 42)
    v1 = (v1 << ((0) & 63))
    if true {
        fmt.Println(v1, v1)
    } else {
        fmt.Println(v2, v1)
    }
    var v17 uint8 = 42
    switch 42 {
    case 1:
        v18 := GenS1{F0: 13}
        _ = v18
    case 300:
        var v19 int64 = -1
        _ = v19
    }
    var v20 int8 = -1
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int = -836
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        v2 = ("foo bar" + ((v2 + (v2 + v2)) + "foo bar"))
        var v24 int64 = 641
        for fi25 := 0; fi25 < 10; fi25++ {
            v17 += (^(v17))
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    if false {
        fmt.Println(v1, v20)
        v26 := GenS2{F0: -1}
        v27 := GenS2{F0: 1}
        _ = v26
        _ = v27
    } else {
        fmt.Println(v2, v17)
        var v28 int64 = 668
        _ = v28
    }
    switch ("hello" + v2) {
    case "\n\t":
        fmt.Println(v20, v20)
    default:
        if (4294967295 <= 2147483647) {
            var v29 int64 = -1
            _ = v29
        } else {
            v1 += v17
            v17 = v17
        }
    }
    v17 = (v1 % ((v17) | 1))
    v2 = v2
    _ = len(v4)
    switch ("foo bar" + "foo bar") {
    case "":
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 int = fi30
            var v32 int32 = -499
            _ = fi30
            _ = v31
            _ = v32
        }
    case "\n\t":
        v2 = "a"
    case "foo bar":
        var v33 bool = true
        _ = v33
    }
    v2 = "hello"
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int = 9223372036854775807
    var v2 int8 = -1
    v3 := []uint32{2147483648}
    var v4 int = v1
    v5 := GenS1{F0: 42}
    var v6 float32 = 1.5
    v7 := GenS1{F0: 32767}
    v1 = 32767
    if (true && (1 == 215)) {
        v8 := GenS2{F0: int8(v1)}
        var v9 int64 = -9223372036854775808
        _ = v8
        _ = v9
    } else {
        var v10 uint16 = 42
        _ = v10
    }
    v4 = 639
    v11 := GenS1{F0: 42}
    _ = len(v3)
    var v12 uint16 = 2
    v1 = v4
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v1, fi13)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi13
    }
    if true {
        v1 -= v1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v3)
    } else {
        v14 := GenS2{F0: (-(v2))}
        _ = v14
    }
    _ = len(v3)
    if true {
        v12 = (v12 << ((18446744073709551615) & 63))
        switch "foo bar" {
        case "\n\t":
            _ = len(v3)
        case "hello":
            if (-104 < 745) {
                var v15 uint = 0
                if (!(2147483648 > 65535)) {
                    v4 = -581
                    v2 = (-(v2))
                    for fi16 := 0; fi16 < 20; fi16++ {
                        if (1 >= 2147483648) {
                            _ = len(v3)
                            if (v12 >= v12) {
                                if false {
                                    v6 = 1.5
                                    var v17 float32 = v6
                                    _ = v17
                                }
                            }
                        }
                        _ = fi16
                    }
                }
                _ = v15
            } else {
                v4 = v1
                fmt.Println(v1, v1, v1)
                v1 = v4
            }
        default:
            var v18 string = ""
            _ = v18
        }
        fmt.Println(v2, v6, v4)
    } else {
        if true {
            if (42 < -1) {
                switch v1 {
                case 0:
                    switch v1 {
                    case -1:
                        var v19 float32 = 3.14159
                        _ = v19
                    }
                default:
                    var v20 int32 = -1
                    _ = v20
                }
            }
        } else {
            var v21 int8 = -41
            if false {
                var v22 uint16 = (v12 ^ 747)
                _ = v22
            } else {
                fmt.Println(v21, v2, v1)
            }
            _ = v21
        }
    }
    switch (v4 / ((v1) | 1)) {
    case -9223372036854775808:
        if len(v3) > 0 {
            v3[0] = 65535
        }
    case 42:
        v2 = -1
    case 127:
        v2 = int8(v12)
    default:
        v4 -= v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

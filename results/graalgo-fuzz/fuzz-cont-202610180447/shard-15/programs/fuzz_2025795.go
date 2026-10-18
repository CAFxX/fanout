package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 32767
    var v3 uint8 = uint8(v2)
    var v4 float32 = 1.5
    v5 := make(map[string]int)
    var v6 uint16 = v2
    var v7 int32 = -731
    switch 447 {
    case -359:
        var v8 string = ((("foo bar" + "") + "hello") + "hello")
        _ = v8
    case 0:
        var v9 uint32 = 0
        _ = v9
    default:
        _ = len(v1)
    }
    v6 = v6
    switch 2147483647 {
    case -9223372036854775808:
        v2 -= v2
    default:
        _ = len(v1)
    }
    switch 194 {
    case 32768:
        if ((v3 >> ((18446744073709551615) & 63)) < (v3 << ((65535) & 63))) {
            if ((0 <= 2) || false) {
                for fi10 := 0; fi10 < 4; fi10++ {
                    if false {
                        v4 = (v4 - 0.1)
                    }
                    v11 := GenS1{F0: 0}
                    v12 := GenS2{F0: 443, F1: true}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 int16 = int16(v2)
                var v14 uint16 = 0
                _ = v13
                _ = v14
            } else {
                var v15 int32 = (v7 << ((1) & 63))
                _ = v15
            }
            switch "foo bar" {
            case "foo bar":
                var v16 uint16 = (^(v6))
                _ = v16
            case "hello":
                var v17 bool = (!(v7 == 1))
                _ = v17
            default:
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS1{F0: 947}
                    v20 := GenS2{F0: 2147483647, F1: true}
                    v21 := []int{2147483647, (fi18 >> ((18446744073709551615) & 63)), fi18, 583}
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
            }
            v6 = (v2 % (((v6 % (((v6 >> ((368) & 63))) | 1))) | 1))
        }
    case 2147483647:
        var v22 uint64 = 979
        _ = v22
    }
    fmt.Println(v4, v2, v7)
    if false {
        if true {
            for fi23 := 0; fi23 < 11; fi23++ {
                v24 := GenS2{F0: 965, F1: true}
                _ = fi23
                _ = v24
            }
            delete(v1, 1)
            fmt.Println(v7, v7)
        }
        fmt.Println(v2, v6, v7)
        if (-1 >= 1) {
            v2 += (v2 - v6)
            v3 = (v3 ^ v3)
            if false {
                v6 = v6
                var v25 string = "hello"
                _ = v25
            }
        }
    }
    if (!(false || (true || (!(float32(((1e10 * 1.5) + 0.0)) <= 1.5))))) {
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 uint = 65535
            for fi28 := 0; fi28 < 3; fi28++ {
                switch 42 {
                case 4294967295:
                    delete(v5, "a")
                default:
                    var v29 int = fi28
                    _ = v29
                }
                fmt.Println(fi28, v4, v2)
                _ = fi28
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 uint32 = 2
                var v32 int32 = (v7 % ((v7) | 1))
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi26
            _ = v27
        }
    } else {
        v6 = (v6 ^ (v2 ^ v6))
    }
    v33 := GenS3{F0: 32767, F1: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}

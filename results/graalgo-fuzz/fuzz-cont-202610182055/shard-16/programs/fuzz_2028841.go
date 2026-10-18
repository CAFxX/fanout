package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
func genFn0() string {
    v7 := GenS1{F0: 42}
    if false {
        var v8 uint64 = 4294967295
        _ = v8
    }
    if ((0.1 - -2.25) <= -2.25) {
        switch -1 {
        case -9223372036854775808:
            switch "hello" {
            case "foo bar":
                var v9 uint32 = 576
                _ = v9
            default:
                v10 := GenS1{F0: 42}
                _ = v10
            }
        case 2147483647:
            var v11 bool = false
            _ = v11
        case 551:
            switch -740 {
            case 1:
                if (1 == -1) {
                    var v12 string = "hello"
                    if false {
                        v13 := []int{813, -9223372036854775808, 422, 32768}
                        v14 := GenS1{F0: 0}
                        v15 := GenS1{F0: 327}
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        if true {
                            if (false || (252 > 128)) {
                                v12 = ("foo bar" + v12)
                                var v16 int32 = -963
                                _ = v16
                            }
                            var v17 int16 = 1000
                            _ = v17
                        }
                        var v18 int32 = -805
                        _ = v18
                    }
                    _ = v12
                } else {
                    var v19 uint32 = 65535
                    v19 = 65535
                    _ = v19
                }
            case 32767:
                var v20 bool = true
                _ = v20
            case -9223372036854775808:
                var v21 uint32 = 65535
                _ = v21
            }
        }
        var v22 string = "\n\t"
        fmt.Println(v22)
        _ = v22
    } else {
        v23 := GenS1{F0: 475}
        v24 := GenS1{F0: 18446744073709551615}
        _ = v23
        _ = v24
    }
    v25 := []int{1, 530, -9223372036854775808}
    _ = v25
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint8, p1 bool, p2 uint) string {
    var v26 string = ("foo bar" + "foo bar")
    p0 = 0
    if (-757 != -32768) {
        var v27 uint64 = 42
        for fi28 := 0; fi28 < 12; fi28++ {
            p0 = p0
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := GenS2{F0: -825}
                fmt.Println(fi29)
                var v31 uint = p2
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
        _ = v27
    }
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return v26
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 uint8 = 127
    var v5 int64 = 9223372036854775807
    var v6 uint8 = (v4 % ((v4) | 1))
    var v32 uint16 = 964
    v6 = (v4 - (v6 | v4))
    var v33 int64 = 1
    v2[4294967295] = false
    fmt.Println(v32, v1)
    var v34 int32 = 2147483647
    var v35 int16 = 529
    v33 += (v33 << ((331) & 63))
    _ = len(v2)
    if false {
        fmt.Println(v32, v35, v32)
        if ((v1 >= (v1 & v1)) || true) {
            for fi36 := 0; fi36 < 7; fi36++ {
                for fi37 := 0; fi37 < 16; fi37++ {
                    var v38 float32 = 1e10
                    _ = fi37
                    _ = v38
                }
                v39 := GenS1{F0: 559}
                v2[(-(fi36))] = true
                _ = fi36
                _ = v39
            }
        }
    } else {
        fmt.Println(v32)
    }
    v40 := GenS2{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}

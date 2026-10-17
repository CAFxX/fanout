package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int32; F1 int16 }
func genFn0() int64 {
    if true {
        var v5 int32 = 720
        _ = v5
    }
    var v6 uint16 = 65535
    _ = v6
    return 1
}

func main() {
    var v1 uint32 = 435
    v2 := make(map[int]uint32)
    var v3 int32 = 42
    v4 := []uint32{(v1 + v1)}
    switch "hello" {
    case "hello":
        fmt.Println(v1, v3)
    case "":
        if true {
            var v7 string = ""
            var v8 uint32 = 65535
            _ = v7
            _ = v8
        }
    case "a":
        v9 := v2[-78]
        _ = v9
    default:
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                if len(v4) > 0 {
                    v4[0] = v1
                }
                _ = fi11
            }
            v12 := GenS1{F0: 90, F1: 18446744073709551615}
            v13 := v2[(fi10 % ((fi10) | 1))]
            _ = fi10
            _ = v12
            _ = v13
        }
    }
    v3 -= v3
    switch -900 {
    case 0:
        delete(v2, 182)
    case 9223372036854775807:
        var v14 bool = false
        _ = v14
    case -847:
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS1{F0: 334, F1: 960}
            if false {
                v1 = 995
            } else {
                if false {
                    v1 += v1
                }
                v2[(fi15 | (fi15 >> ((558) & 63)))] = 1
            }
            if (!true) {
                v17 := GenS1{F0: 42, F1: 186}
                _ = v17
            }
            _ = fi15
            _ = v16
        }
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 19; fi19++ {
            v1 += v1
            _ = fi19
        }
        v3 = v3
        _ = fi18
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        if len(v4) > 0 {
            v4[0] = 65535
        }
        if len(v4) > 0 {
            v4[0] = (^(v1))
        }
        _ = fi20
    }
    v21 := GenS2{F0: 622, F1: -479}
    switch ((("\n\t" + ("hello" + "")) + ("\n\t" + "foo bar")) + "hello") {
    case "hello":
        _ = len(v2)
    case "":
        if (!false) {
            switch ("foo bar" + "foo bar") {
            case "a":
                if false {
                    var v22 uint = 0
                    _ = v22
                } else {
                    switch -322 {
                    case 42:
                        v1 = (v1 << ((476) & 63))
                    case 4294967295:
                        if false {
                            var v23 bool = (true || false)
                            _ = v23
                        }
                    }
                    for fi24 := 0; fi24 < 4; fi24++ {
                        delete(v2, (fi24 & (fi24 - (fi24 - 774))))
                        v3 = 42
                        fmt.Println(v1, v3, v1)
                        _ = fi24
                    }
                }
            case "\n\t":
                var v25 int8 = -77
                _ = v25
            case "foo bar":
                var v26 uint16 = 65535
                _ = v26
            default:
                switch "hello" {
                case "\n\t":
                    for fi27 := 0; fi27 < 11; fi27++ {
                        delete(v2, -921)
                        fmt.Println(v3)
                        _ = fi27
                    }
                case "a":
                    v3 += -1
                case "foo bar":
                    v28 := GenS3{F0: v3, F1: 1000}
                    _ = v28
                }
            }
        } else {
            var v29 int8 = -127
            var v30 int16 = 1
            _ = v29
            _ = v30
        }
    case "\n\t":
        v3 = (^(v3))
    }
    var v31 int64 = -770
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
}

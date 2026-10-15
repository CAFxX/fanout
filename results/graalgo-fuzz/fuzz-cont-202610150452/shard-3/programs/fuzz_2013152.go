package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = 42
    v3 := []string{}
    var v4 uint64 = 0
    var v5 uint16 = 42
    switch 902 {
    case -297:
        if false {
            var v6 uint32 = 65535
            switch 4294967295 {
            case -872:
                delete(v1, 519)
            case 32767:
                for fi7 := 0; fi7 < 7; fi7++ {
                    v5 = 2
                    _ = fi7
                }
            }
            _ = v6
        } else {
            if ("a" == ("foo bar" + ("hello" + ("\n\t" + "\n\t")))) {
                v8 := GenS1{F0: (v4 << ((905) & 63)), F1: (v2 ^ (v2 * 800))}
                fmt.Println(v2)
                v5 += 1
                _ = v8
            } else {
                for fi9 := 0; fi9 < 10; fi9++ {
                    switch "hello" {
                    case "":
                        v4 = v4
                    }
                    var v10 int16 = -1
                    _ = fi9
                    _ = v10
                }
                v2 = 168
            }
            v5 = (v5 + (v5 & v5))
            v5 = v5
        }
    case 2147483647:
        fmt.Println(v4, v5, v5)
    default:
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = (v5 & 403)
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 uint = 2
                v2 = -630
                switch fi13 {
                case 4294967295:
                    v2 += 25
                case 32768:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -9223372036854775808:
                    var v15 int16 = 32767
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            v12 -= (v12 | v12)
            _ = fi11
            _ = v12
        }
    }
    v4 -= 93
    var v16 int8 = -43
    v16 -= (v16 << ((255) & 63))
    if false {
        for fi17 := 0; fi17 < 19; fi17++ {
            if ((float64(-2.25) + float64(float32(1.5))) == 1.5) {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v4 = (v4 - 965)
                    v4 = v4
                    _ = fi18
                }
            }
            _ = fi17
        }
    }
    fmt.Println(v5)
    delete(v1, 32767)
    v19 := GenS2{F0: 4294967295}
    for fi20 := 0; fi20 < 8; fi20++ {
        v2 = 177
        v5 -= v5
        _ = fi20
    }
    var v21 uint16 = v5
    if false {
        for fi22 := 0; fi22 < 10; fi22++ {
            delete(v1, -8)
            _ = fi22
        }
        var v23 int32 = (v2 << ((0) & 63))
        _ = v23
    } else {
        delete(v1, -801)
    }
    if true {
        v1[127] = (false || (42 != -9223372036854775808))
    }
    v16 = 0
    var v24 int32 = v2
    var v25 uint = 42
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}

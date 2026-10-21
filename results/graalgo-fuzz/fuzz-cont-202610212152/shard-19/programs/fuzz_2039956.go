package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 bool }
func genFn0(p0 uint64) string {
    if true {
        switch "a" {
        case "\n\t":
            if true {
                v5 := []int{}
                v6 := GenS1{F0: "\n\t"}
                _ = v5
                _ = v6
            } else {
                switch 2147483647 {
                case -68:
                    var v7 string = "foo bar"
                    _ = v7
                case 693:
                    v8 := GenS1{F0: ("a" + ("hello" + ("a" + ("\n\t" + (("" + "") + "foo bar")))))}
                    _ = v8
                case -9223372036854775808:
                    p0 = (p0 * 0)
                }
            }
        case "":
            v9 := []int{}
            _ = v9
        case "hello":
            var v10 bool = false
            _ = v10
        }
    }
    switch 127 {
    case 127:
        var v11 int64 = 0
        _ = v11
    case 42:
        p0 += 275
    case 0:
        if (32767 <= 0) {
            var v12 uint = 2
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint32 = 4294967295
                switch fi13 {
                case 128:
                    v15 := GenS1{F0: ""}
                    _ = v15
                default:
                    if ((!(0 >= int16(v12))) && false) {
                        fmt.Println(p0, v12, fi13)
                        v16 := GenS2{F0: "\n\t", F1: 127}
                        v14 -= v14
                        _ = v16
                    }
                }
                var v17 int32 = 2147483647
                _ = fi13
                _ = v14
                _ = v17
            }
            _ = v12
        }
    default:
        p0 -= (p0 << ((498) & 63))
    }
    _ = p0
    return ("" + "\n\t")
}

func main() {
    var v1 uint64 = 285
    var v2 int = -487
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v18 bool = false
    switch "foo bar" {
    case "":
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS1{F0: "\n\t"}
            var v21 int64 = 1
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if false {
        var v22 uint8 = 1
        _ = v22
    }
    var v23 uint32 = 2
    if (!(!(!(true || v18)))) {
        v18 = v18
    }
    v1 -= (v1 | (v1 ^ (v1 % (((^(v1))) | 1))))
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS3{F0: 42, F1: false}
        fmt.Println(v1, v18)
        v1 = 4294967295
        _ = fi24
        _ = v25
    }
    v23 += (v23 ^ 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}

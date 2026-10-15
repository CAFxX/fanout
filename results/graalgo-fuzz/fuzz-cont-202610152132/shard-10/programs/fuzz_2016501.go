package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
func genFn0() int8 {
    if (!false) {
        v6 := []int{1}
        var v7 int16 = 753
        var v8 uint16 = 65535
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 uint16 = 0
        v9 = (v9 & v9)
        _ = v9
    }
    var v10 string = "hello"
    v11 := GenS1{F0: (32767 > 32767)}
    _ = v10
    _ = v11
    return 0
}
func genFn1() bool {
    var v12 string = ("a" + "a")
    v12 = "a"
    v13 := []int{389, 0, 32768}
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 uint64 = 2
    var v2 uint = 255
    v3 := make(map[int]int)
    var v4 int32 = 969
    v5 := make(map[int]int)
    v4 += -1
    for fi14 := 0; fi14 < 12; fi14++ {
        v2 = (v2 >> (((^(v2))) & 63))
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    switch ("hello" + ((((("" + "foo bar") + "\n\t") + "\n\t") + "hello") + "hello")) {
    case "hello":
        var v16 int8 = 100
        _ = v16
    case "\n\t":
        var v17 uint64 = (^(v1))
        _ = v17
    case "a":
        fmt.Println(v4)
    default:
        for fi18 := 0; fi18 < 12; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                v4 -= -308
                _ = fi19
            }
            _ = fi18
        }
    }
    var v20 int32 = (v4 % ((2147483647) | 1))
    delete(v5, 127)
    v20 = v4
    v20 = v4
    switch 9223372036854775807 {
    case -9223372036854775808:
        for fi21 := 0; fi21 < 13; fi21++ {
            if false {
                fmt.Println(v1, fi21)
                for fi22 := 0; fi22 < 5; fi22++ {
                    delete(v3, fi21)
                    fmt.Println(v4, v2)
                    v23 := GenS2{F0: 431}
                    _ = fi22
                    _ = v23
                }
                v24 := GenS3{F0: 65535}
                _ = v24
            }
            if true {
                v4 = (-(v20))
            } else {
                v4 += (v20 / ((v20) | 1))
            }
            _ = fi21
        }
    case 32767:
        fmt.Println(v1, v4, v20)
    default:
        var v25 uint8 = 86
        _ = v25
    }
    var v26 uint = (v2 / ((255) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int16; F1 uint; F2 int16 }
func genFn0() int16 {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint16 = 65535
        v6 -= v6
        _ = fi5
        _ = v6
    }
    switch "foo bar" {
    case "a":
        v7 := GenS1{F0: 775, F1: 161}
        _ = v7
    }
    var v8 int16 = -1
    v9 := GenS1{F0: 4294967295, F1: 2147483647}
    _ = v8
    _ = v9
    return (v8 | v8)
}
func genFn1() int8 {
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(fi10, fi11)
                _ = fi11
            }
            var v12 int32 = -2147483648
            var v13 float32 = 1.5
            _ = fi10
            _ = v12
            _ = v13
        }
    }
    v14 := GenS2{F0: "hello", F1: 32767}
    _ = v14
    return -128
}

func main() {
    var v1 float64 = -2.25
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v4 int8 = 0
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            delete(v3, 65535)
            _ = fi16
        }
        v17 := GenS2{F0: "", F1: 32767}
        v18 := []int{32767, (fi15 % ((fi15) | 1)), fi15}
        _ = fi15
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := GenS2{F0: "hello", F1: 1}
        v3[2] = (("foo bar" + "hello") + "\n\t")
        v21 := GenS3{F0: 32767, F1: 18446744073709551615, F2: 735}
        _ = fi19
        _ = v20
        _ = v21
    }
    v4 = genFn1()
    var v22 uint64 = 42
    v1 = 1e10
    var v23 int32 = v2
    switch "a" {
    case "":
        v22 = (v22 ^ 233)
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        v23 += v23
    default:
        fmt.Println(v2)
    }
    switch ("\n\t" + ((("\n\t" + "hello") + "a") + ("a" + (("foo bar" + "foo bar") + "hello")))) {
    case "a":
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 19; fi25++ {
                fmt.Println(fi24, v1)
                var v26 int8 = genFn1()
                _ = fi25
                _ = v26
            }
            var v27 int64 = 0
            v23 -= (v23 * -2147483648)
            _ = fi24
            _ = v27
        }
    case "hello":
        v23 = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}

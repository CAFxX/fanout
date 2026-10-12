package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int16; F1 int; F2 string }
func genFn0(p0 int8, p1 int8, p2 uint) uint32 {
    if (true && (-9223372036854775808 != -9223372036854775808)) {
        fmt.Println(p0, p2, p0)
        v5 := GenS1{F0: ((p2 % ((612) | 1)) == p2)}
        v6 := []int{42, 1}
        _ = v5
        _ = v6
    }
    p2 = p2
    v7 := []int{-9223372036854775808, 42, 243}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}
func genFn1(p0 int64, p1 int) uint32 {
    v8 := []int{p1, p1}
    if true {
        var v9 uint16 = 65535
        switch "hello" {
        case "\n\t":
            for fi10 := 0; fi10 < 10; fi10++ {
                if (0 < 337) {
                    switch 953 {
                    case 0:
                        fmt.Println(v9)
                    case 42:
                        v11 := GenS2{F0: 255, F1: 2147483648}
                        _ = v11
                    default:
                        v9 = v9
                    }
                } else {
                    v8 = append(v8, fi10)
                }
                for fi12 := 0; fi12 < 15; fi12++ {
                    v13 := GenS1{F0: false}
                    _ = fi12
                    _ = v13
                }
                _ = fi10
            }
        }
        _ = v9
    }
    var v14 int16 = -32768
    _ = p0
    _ = p1
    _ = v14
    _ = v8
    return 2147483647
}

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]uint32)
    v3 := []string{("a" + ("foo bar" + "hello")), (("" + "foo bar") + "hello"), "", ("\n\t" + ("hello" + (("hello" + ((("\n\t" + "\n\t") + "hello") + "foo bar")) + "\n\t")))}
    v4 := make(map[int]string)
    if ((4294967295 == 4294967295) && true) {
        delete(v4, 42)
    }
    v1 = v1
    var v15 uint8 = 149
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 := GenS2{F0: 42, F1: 1}
    var v17 int8 = 31
    for fi18 := 0; fi18 < 4; fi18++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi18
    }
    var v19 uint8 = (^(v15))
    v1 = 57
    var v20 string = ("foo bar" + "foo bar")
    v19 = v19
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint = 255
        fmt.Println(v20)
        _ = fi21
        _ = v22
    }
    v19 -= (v15 * (v19 / ((v15) | 1)))
    for fi23 := 0; fi23 < 2; fi23++ {
        v19 = (v19 & v15)
        if (!false) {
            fmt.Println(v17)
        } else {
            v24 := GenS3{F0: 32767, F1: fi23, F2: v20}
            fmt.Println(v20, v1)
            v25 := v2[genFn1(1, fi23)]
            _ = v24
            _ = v25
        }
        _ = fi23
    }
    switch v20 {
    case "foo bar":
        var v26 float32 = (float32(1.5) / float32((0.1 * 1e10)))
        _ = v26
    case "\n\t":
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint8 = 46
            v20 = "\n\t"
            _ = fi27
            _ = v28
        }
    case "a":
        var v29 string = ""
        _ = v29
    }
    var v30 int = -231
    if false {
        var v31 int64 = 1
        v17 += (^(v17))
        v19 = 10
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int; F1 uint }
func genFn0(p0 int32, p1 string, p2 uint64) int32 {
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 32767}
        if true {
            if false {
                for fi7 := 0; fi7 < 4; fi7++ {
                    p2 -= (p2 + p2)
                    switch (("\n\t" + p1) + p1) {
                    case "a":
                        fmt.Println(p2)
                    case "\n\t":
                        p1 = ("foo bar" + p1)
                    case "":
                        var v8 uint64 = (p2 ^ (p2 % ((0) | 1)))
                        v9 := []int{fi5}
                        _ = v8
                        _ = v9
                    }
                    v10 := GenS2{F0: p1, F1: 0}
                    _ = fi7
                    _ = v10
                }
                var v11 int64 = -1
                _ = v11
            }
            for fi12 := 0; fi12 < 7; fi12++ {
                v13 := []int{}
                var v14 uint16 = 18
                v14 = (v14 | (v14 >> ((633) & 63)))
                _ = fi12
                _ = v13
                _ = v14
            }
            if true {
                switch 127 {
                case 42:
                    fmt.Println(p0, fi5, fi5)
                }
                var v15 int64 = 382
                v16 := GenS3{F0: 2, F1: -9223372036854775808}
                _ = v15
                _ = v16
            }
        }
        p0 -= -181
        _ = fi5
        _ = v6
    }
    v17 := []int{42, 128, -682, -659}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return p0
}
func genFn1(p0 uint32, p1 uint8) uint64 {
    for fi18 := 0; fi18 < 18; fi18++ {
        p1 += p1
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := []int{fi18, fi19, (^(fi19)), (^(fi19))}
            var v21 float32 = (0.1 - 3.14159)
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    var v22 string = "foo bar"
    fmt.Println(p0, v22, v22)
    p1 -= p1
    v23 := []int{208, 32767, 4294967295}
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 127
    var v3 bool = false
    var v4 uint8 = 42
    var v24 uint = 255
    if v3 {
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 uint16 = 65535
            v27 := GenS4{F0: (fi25 % ((v2) | 1)), F1: v24}
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 int16 = -32768
                _ = fi28
                _ = v29
            }
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    _ = len(v1)
    if v3 {
        v4 -= 2
    } else {
        var v30 int16 = -14
        var v31 int = v2
        _ = v30
        _ = v31
    }
    v32 := GenS4{F0: v2, F1: v24}
    v3 = false
    var v33 uint = v24
    var v34 string = ("" + (("a" + (("hello" + "") + "")) + ("" + "foo bar")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}

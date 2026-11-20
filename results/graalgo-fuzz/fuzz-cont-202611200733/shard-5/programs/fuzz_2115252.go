package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0() int {
    var v5 uint = 42
    v5 += (^(v5))
    fmt.Println(v5)
    _ = v5
    return -420
}
func genFn1() int32 {
    var v6 int16 = 32767
    v7 := []int{32767}
    if len(v7) > 0 {
        v7[0] = 127
    }
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 uint64 = 42
    switch genFn0() {
    case -295:
        var v8 int32 = -799
        _ = v8
    }
    var v9 uint64 = (^(v4))
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint64 = (v4 - v9)
        var v12 int16 = 694
        _ = fi10
        _ = v11
        _ = v12
    }
    v4 = (v4 | (v4 >> ((461) & 63)))
    v4 -= 862
    switch "hello" {
    case "":
        if (2 >= 42) {
            fmt.Println(v4)
        } else {
            v4 = (v4 + v4)
            var v13 string = ""
            v14 := []int{v1, (v1 + (v1 ^ v1)), v1}
            _ = v13
            _ = v14
        }
    default:
        var v15 int32 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS1{F0: 42, F1: 381, F2: 42}
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: 857, F1: 255, F2: 255}
    if true {
        v19 := GenS1{F0: 0, F1: 463, F2: 685}
        var v20 uint16 = 0
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := []int{2147483647, v1, (v1 + 527), (v1 & genFn0())}
            _ = fi21
            _ = v22
        }
        v9 = 590
        if false {
            v4 = v4
        } else {
            v23 := GenS1{F0: 42, F1: 2, F2: 1}
            _ = v23
        }
    }
    delete(v3, 1)
    if true {
        switch 123 {
        case -939:
            if (396 == -1) {
                if true {
                    var v24 int64 = 9223372036854775807
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        var v25 uint64 = v9
                        _ = v25
                    case "foo bar":
                        delete(v3, uint32(v4))
                    case "hello":
                        var v26 string = "a"
                        _ = v26
                    default:
                        for fi27 := 0; fi27 < 11; fi27++ {
                            var v28 uint8 = 255
                            _ = fi27
                            _ = v28
                        }
                    }
                    switch (("\n\t" + ("" + "a")) + "\n\t") {
                    case "":
                        var v29 uint16 = 65535
                        _ = v29
                    case "hello":
                        var v30 string = ("hello" + ("foo bar" + ""))
                        _ = v30
                    case "a":
                        for fi31 := 0; fi31 < 20; fi31++ {
                            var v32 float32 = float32(0.1)
                            var v33 int = (-(v1))
                            fmt.Println(v33, v4, v33)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                    _ = v24
                } else {
                    var v34 uint32 = 695
                    _ = v34
                }
            }
        case 9223372036854775807:
            v9 += (v4 & 4294967295)
        }
    }
    if false {
        var v35 bool = false
        switch "\n\t" {
        case "foo bar":
            for fi36 := 0; fi36 < 6; fi36++ {
                if true {
                    v4 -= 92
                    switch fi36 {
                    case -2147483648:
                        v35 = v35
                    case -749:
                        v4 -= (v9 * v9)
                    }
                    fmt.Println(v9)
                }
                for fi37 := 0; fi37 < 6; fi37++ {
                    v35 = (32767 > 255)
                    _ = fi37
                }
                switch "a" {
                case "foo bar":
                    var v38 int32 = 392
                    _ = v38
                case "hello":
                    var v39 bool = false
                    _ = v39
                case "a":
                    v1 = 0
                }
                _ = fi36
            }
        default:
            for fi40 := 0; fi40 < 1; fi40++ {
                fmt.Println(v35)
                v35 = false
                _ = fi40
            }
        }
        _ = v35
    }
    v41 := GenS2{F0: v1, F1: 2}
    v42 := GenS2{F0: 42, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v9
}

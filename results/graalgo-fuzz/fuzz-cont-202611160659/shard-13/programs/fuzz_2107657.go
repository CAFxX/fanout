package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int32 }
func genFn0() uint8 {
    var v5 uint32 = 57
    v5 = v5
    var v6 float64 = 0.1
    if true {
        if true {
            if true {
                v7 := []int{127, 42, 32768}
                v8 := GenS1{F0: -1}
                if ((v5 ^ (v5 * 4294967295)) > v5) {
                    if false {
                        fmt.Println(v6)
                        var v9 uint16 = 1
                        fmt.Println(v5, v9, v5)
                        _ = v9
                    } else {
                        var v10 int16 = 0
                        v10 -= -32768
                        _ = v10
                    }
                }
                _ = v7
                _ = v8
            }
            var v11 float64 = v6
            v6 = 0.0
            _ = v11
        }
    } else {
        v12 := []int{-260}
        v6 -= (3.14159 - v6)
        v5 = v5
        _ = v12
    }
    var v13 float32 = 0.0
    _ = v13
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 bool = (false && true)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    v4 := []uint32{937, 42, 4294967295}
    fmt.Println(v1)
    v1 = false
    v14 := GenS1{F0: -128}
    var v15 uint64 = 1
    var v16 uint8 = 255
    for fi17 := 0; fi17 < 20; fi17++ {
        if false {
            if true {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v15 -= 42
                    var v19 uint = 18446744073709551615
                    _ = fi18
                    _ = v19
                }
                v20 := GenS2{F0: 4294967295, F1: 127, F2: 42}
                _ = v20
            } else {
                if ((fi17 >> ((570) & 63)) != fi17) {
                    var v21 uint16 = 828
                    _ = v21
                }
            }
            var v22 string = "hello"
            _ = v22
        }
        if (0 > 1) {
            v23 := v2[fi17]
            if v1 {
                v24 := GenS2{F0: 42, F1: (v16 - v16), F2: 2147483647}
                _ = v24
            }
            var v25 int8 = int8(fi17)
            _ = v23
            _ = v25
        } else {
            if v1 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v1 = ((1000 > int16(fi17)) && v1)
            }
            if (v1 && (!(1e10 <= float32((1.5 + float64((((1.5 * 0.1) / -2.25) / 1e10))))))) {
                switch -153 {
                case 128:
                    var v26 uint = 65535
                    _ = v26
                case 543:
                    var v27 int64 = 1
                    _ = v27
                case 127:
                    fmt.Println(v1, v15, v15)
                default:
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 uint = 4294967295
                        if len(v4) > 0 {
                            v4[0] = 416
                        }
                        var v30 int32 = 376
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := GenS1{F0: 3}
                    var v33 uint16 = 255
                    var v34 uint = 4294967295
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            } else {
                var v35 uint = 0
                fmt.Println(fi17, v1, v1)
                var v36 string = (("\n\t" + (("foo bar" + "\n\t") + "foo bar")) + (("hello" + ("" + ("hello" + ""))) + "hello"))
                _ = v35
                _ = v36
            }
        }
        _ = fi17
    }
    var v37 int8 = -65
    var v38 uint = 15
    delete(v3, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}

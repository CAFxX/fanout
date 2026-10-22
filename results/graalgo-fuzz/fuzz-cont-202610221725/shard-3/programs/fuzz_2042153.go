package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
type GenS4 struct { F0 string }
func genFn0(p0 uint32, p1 uint16, p2 uint) bool {
    var v7 float32 = -2.25
    var v8 uint = 98
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := GenS1{F0: 0, F1: 42}
            if true {
                switch ("hello" + "a") {
                case "\n\t":
                    fmt.Println(fi9, p0)
                }
            }
            _ = fi10
            _ = v11
        }
        if false {
            v12 := []int{fi9, fi9, -2147483648}
            _ = v12
        } else {
            v13 := []int{fi9, -1}
            for fi14 := 0; fi14 < 3; fi14++ {
                switch "\n\t" {
                case "":
                    var v15 int8 = 0
                    _ = v15
                case "a":
                    fmt.Println(p0)
                case "\n\t":
                    v16 := GenS2{F0: 251}
                    _ = v16
                default:
                    var v17 int8 = 24
                    _ = v17
                }
                _ = fi14
            }
            var v18 float32 = float32(-2.25)
            _ = v13
            _ = v18
        }
        p0 = 590
        _ = fi9
    }
    v8 = (p2 & 318)
    fmt.Println(p1, v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint16 = 255
    var v2 float32 = -2.25
    v3 := make(map[int]bool)
    var v4 int8 = 127
    v5 := []uint32{}
    var v6 float32 = v2
    v1 -= 2
    v4 = -1
    v19 := v3[128]
    v20 := GenS3{F0: ("" + (("\n\t" + (("foo bar" + ("hello" + "a")) + "hello")) + "foo bar")), F1: -32768, F2: v4}
    v21 := GenS4{F0: "hello"}
    v22 := v3[-656]
    if v22 {
        for fi23 := 0; fi23 < 10; fi23++ {
            v22 = v19
            for fi24 := 0; fi24 < 7; fi24++ {
                var v25 uint32 = 1
                _ = fi24
                _ = v25
            }
            var v26 string = ""
            _ = fi23
            _ = v26
        }
        switch -867 {
        case 0:
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 uint64 = 2
                _ = fi27
                _ = v28
            }
        case -617:
            var v29 int64 = 0
            _ = v29
        }
        var v30 int32 = -581
        _ = v30
    }
    if len(v5) > 0 {
        v5[0] = 548
    }
    v2 = v2
    if v19 {
        if v19 {
            for fi31 := 0; fi31 < 1; fi31++ {
                switch "hello" {
                case "\n\t":
                    for fi32 := 0; fi32 < 12; fi32++ {
                        if len(v5) > 0 {
                            v5[0] = 735
                        }
                        v2 += (v6 * 0.1)
                        v6 = (v2 / -2.25)
                        if (9223372036854775807 != -910) {
                            _ = len(v5)
                        }
                        _ = fi32
                    }
                default:
                    var v33 int = fi31
                    _ = v33
                }
                for fi34 := 0; fi34 < 11; fi34++ {
                    v19 = v22
                    v35 := GenS4{F0: "hello"}
                    var v36 bool = (2 <= 4294967295)
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                v22 = false
                _ = fi31
            }
            var v37 int32 = -992
            _ = v37
        } else {
            var v38 int8 = v4
            var v39 uint = 255
            var v40 int64 = 0
            _ = v38
            _ = v39
            _ = v40
        }
        if (true && (v22 || (!v19))) {
            v1 = 2
            var v41 uint16 = v1
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v41
        } else {
            v22 = (((0.0 * 0.0) / 0.1) != 0.0)
        }
        v6 += (3.14159 - 0.0)
    } else {
        delete(v3, 128)
    }
    var v42 int = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}

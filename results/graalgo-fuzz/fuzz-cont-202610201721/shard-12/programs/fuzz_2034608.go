package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint32) uint32 {
    var v4 int8 = 39
    v4 = 0
    _ = p0
    _ = v4
    return (p0 << ((1) & 63))
}
func genFn1(p0 uint8, p1 uint64) uint16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint = 42
        v7 := []int{fi5, (fi5 | (-(fi5))), 4294967295, (-(fi5))}
        _ = len(v7)
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p0, p0)
    v8 := []int{-9223372036854775808, 4294967295, -666, -906}
    _ = p0
    _ = p1
    _ = v8
    return 32767
}

func main() {
    var v1 uint8 = 11
    var v2 float32 = 0.1
    var v3 uint32 = 2147483647
    switch ("" + "\n\t") {
    case "\n\t":
        switch 128 {
        case 9223372036854775807:
            switch 128 {
            case -1:
                if (255 < 687) {
                    for fi9 := 0; fi9 < 14; fi9++ {
                        v1 -= (^(v1))
                        for fi10 := 0; fi10 < 15; fi10++ {
                            fmt.Println(fi10, fi10)
                            _ = fi10
                        }
                        _ = fi9
                    }
                } else {
                    var v11 int32 = -46
                    if (v1 != v1) {
                        v11 = 0
                        var v12 uint32 = v3
                        _ = v12
                    }
                    v13 := GenS1{F0: (-(v11))}
                    _ = v11
                    _ = v13
                }
            case -994:
                v14 := GenS1{F0: 2147483647}
                _ = v14
            }
        }
    default:
        var v15 int16 = -32768
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v3)
        fmt.Println(v3, fi16, v2)
        v17 := []int{(fi16 << ((4294967295) & 63))}
        _ = fi16
        _ = v17
    }
    v2 = 1.5
    v18 := GenS1{F0: -1}
    v19 := []int{-978, -555, 32767}
    var v20 uint8 = 0
    v19 = append(v19, 42)
    v2 -= v2
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            if false {
                v1 = (v20 / ((133) | 1))
            } else {
                var v23 float32 = v2
                var v24 int8 = -20
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        for fi25 := 0; fi25 < 19; fi25++ {
            v3 += (v3 - v3)
            _ = fi25
        }
        _ = fi21
    }
    v2 -= v2
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 bool = (true && false)
        switch (fi26 ^ (^(fi26))) {
        case -924:
            v3 = (v3 & v3)
        }
        _ = len(v19)
        _ = fi26
        _ = v27
    }
    var v28 int8 = 1
    if true {
        if (32767 <= -875) {
            for fi29 := 0; fi29 < 2; fi29++ {
                fmt.Println(v3)
                fmt.Println(v20)
                _ = fi29
            }
            for fi30 := 0; fi30 < 14; fi30++ {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                var v31 uint8 = 42
                _ = fi30
                _ = v31
            }
        }
        var v32 int = 1
        _ = v32
    }
    v1 = (v20 ^ (v20 / ((v1) | 1)))
    switch "a" {
    case "hello":
        v3 = (v3 & v3)
    default:
        switch -2147483648 {
        case 969:
            fmt.Println(v2)
        default:
            var v33 uint8 = v1
            _ = v33
        }
    }
    if false {
        v2 = v2
        if len(v19) > 0 {
            v19[0] = -2147483648
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
}

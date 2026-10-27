package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint64, p1 int32, p2 int32) bool {
    fmt.Println(p1)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int64 = -482
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 int64 = 694
    v2 := make(map[string]string)
    var v3 uint8 = 1
    var v4 int32 = -449
    var v5 float64 = 1e10
    var v6 int16 = 42
    var v9 int16 = (-(v6))
    v9 = v9
    v1 += (v1 - v1)
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v4, v5, v5)
        v9 = v9
        switch ("a" + "a") {
        case "a":
            v6 = 1000
        case "":
            v11 := []int{}
            _ = v11
        default:
            var v12 bool = false
            _ = v12
        }
        _ = fi10
    }
    var v13 int16 = v9
    switch 32768 {
    case -354:
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := v2["foo bar"]
            _ = fi14
            _ = v15
        }
    case -758:
        fmt.Println(v5)
    case 702:
        if (2147483648 < 4294967295) {
            fmt.Println(v5, v3)
            if false {
                if false {
                    v3 -= (^(v3))
                }
                var v16 uint64 = 42
                _ = v16
            } else {
                v6 += v13
            }
        }
    default:
        v17 := GenS1{F0: 4294967295}
        _ = v17
    }
    v6 += (v9 + 0)
    switch ("hello" + (("foo bar" + ("" + "a")) + ("foo bar" + "hello"))) {
    case "\n\t":
        delete(v2, "")
    default:
        v5 = v5
    }
    var v18 float32 = ((0.0 - (1e10 / float32(1e10))) * -2.25)
    fmt.Println(v9)
    v19 := GenS1{F0: 1}
    var v20 string = ""
    v18 = v18
    var v21 float64 = v5
    if ((v4 ^ v4) >= (^(v4))) {
        switch ("foo bar" + "hello") {
        case "\n\t":
            var v22 string = "a"
            _ = v22
        case "":
            if true {
                v23 := GenS1{F0: 18446744073709551615}
                var v24 uint16 = 90
                _ = v23
                _ = v24
            } else {
                if true {
                    fmt.Println(v9)
                    for fi25 := 0; fi25 < 5; fi25++ {
                        var v26 bool = false
                        var v27 uint64 = 645
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                    v6 = (v9 & v13)
                }
                v13 = (v9 & (v6 << ((65535) & 63)))
                for fi28 := 0; fi28 < 6; fi28++ {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v30 := v2[v20]
                        v31 := []int{fi28, fi28, fi29, 4294967295}
                        v5 = (v21 / float64(((0.0 + ((v18 + (((v18 - (1.5 * 0.1)) / (((((((0.1 / ((v18 / (v18 - (-2.25 + ((float32((v5 / 1.0)) + (((float32(-2.25) * v18) - (v18 + 0.1)) * 3.14159)) / (v18 + v18))))) - v18)) * ((v18 - v18) + 1.5)) * v18) / (v18 - v18)) / float32(v5)) * float32(0.1)) - float32(((0.0 - v21) + float64((v18 / v18)))))) - (1.5 - v18))) / ((float32(v21) * v18) - (0.1 / v18)))) / v18)))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = len(v2)
                    _ = fi28
                }
            }
        case "hello":
            var v32 float64 = (v21 / v5)
            _ = v32
        default:
            var v33 int32 = (v4 * 2147483647)
            _ = v33
        }
        v20 = "a"
    } else {
        fmt.Println(v13, v3)
        v21 = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

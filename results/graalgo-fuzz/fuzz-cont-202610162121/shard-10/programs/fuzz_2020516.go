package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0(p0 int64, p1 string) string {
    p1 = p1
    p0 = (p0 | (p0 / ((-498) | 1)))
    for fi5 := 0; fi5 < 3; fi5++ {
        if false {
            switch "hello" {
            case "":
                if (fi5 >= fi5) {
                    for fi6 := 0; fi6 < 20; fi6++ {
                        v7 := []int{(fi5 - 42), fi6, 9223372036854775807}
                        var v8 bool = false
                        fmt.Println(v8, p0, fi6)
                        _ = fi6
                        _ = v7
                        _ = v8
                    }
                    var v9 int64 = p0
                    fmt.Println(p0, p0, p1)
                    _ = v9
                }
            case "a":
                for fi10 := 0; fi10 < 17; fi10++ {
                    p1 = (p1 + p1)
                    _ = fi10
                }
            }
        } else {
            switch p1 {
            case "\n\t":
                v11 := GenS1{F0: 2, F1: 0}
                _ = v11
            default:
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := []int{-2147483648}
                    switch ("foo bar" + ((((p1 + p1) + ("foo bar" + ("hello" + p1))) + (p1 + "\n\t")) + ((p1 + p1) + "hello"))) {
                    case "hello":
                        v14 := GenS2{F0: 1, F1: 18446744073709551615, F2: 1}
                        p0 = (p0 + 9223372036854775807)
                        p1 = ((p1 + p1) + "a")
                        _ = v14
                    case "foo bar":
                        p1 = "hello"
                    case "":
                        var v15 uint64 = 1
                        _ = v15
                    default:
                        p1 = "\n\t"
                    }
                    _ = fi12
                    _ = v13
                }
            }
            p0 = p0
            p0 += -9223372036854775808
        }
        p0 += 1
        _ = fi5
    }
    v16 := []int{42, 1, -931, 1}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v16
    return p1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{4294967295}
    v3 := []uint32{}
    var v4 string = ""
    switch 32768 {
    case 127:
        if ((!((false || (814 <= 500)) && false)) || true) {
            var v17 uint16 = 700
            _ = v17
        }
    }
    v4 = "a"
    fmt.Println(v4)
    v18 := GenS2{F0: 320, F1: 0, F2: 683}
    v4 = v4
    var v19 int64 = 0
    v20 := GenS2{F0: 2147483648, F1: 144, F2: 18446744073709551615}
    v21 := GenS1{F0: 0, F1: 177}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint8, p1 int, p2 uint64) uint64 {
    if false {
        v4 := GenS1{F0: 0}
        p0 = (p0 >> ((286) & 63))
        _ = v4
    } else {
        v5 := []int{(p1 >> ((2) & 63)), p1, (p1 | (p1 * -2147483648))}
        p0 = (p0 * (p0 << ((2) & 63)))
        _ = v5
    }
    var v6 bool = false
    p0 += (^(p0))
    var v7 uint32 = 696
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    var v1 uint64 = 69
    var v2 bool = (-80 > 2147483647)
    var v3 uint8 = 22
    fmt.Println(v1)
    for fi8 := 0; fi8 < 5; fi8++ {
        switch ("a" + "\n\t") {
        case "a":
            fmt.Println(v3, v2, v3)
        case "\n\t":
            v9 := GenS1{F0: 18446744073709551615}
            _ = v9
        case "":
            for fi10 := 0; fi10 < 7; fi10++ {
                if (false || v2) {
                    v11 := GenS1{F0: 18446744073709551615}
                    _ = v11
                } else {
                    if (1.5 != -2.25) {
                        v1 = v1
                        v2 = v2
                        var v12 uint32 = 0
                        _ = v12
                    }
                    fmt.Println(v3, fi10, v1)
                }
                var v13 int16 = 32767
                _ = fi10
                _ = v13
            }
        default:
            fmt.Println(v1, v1, v3)
        }
        _ = fi8
    }
    if (883 >= 0) {
        var v14 uint8 = (^(v3))
        v14 += v3
        _ = v14
    }
    var v15 uint32 = 2147483647
    v3 = (^(v3))
    v3 += (^(v3))
    var v16 uint = 42
    v3 += 2
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS1{F0: v16}
        _ = fi17
        _ = v18
    }
    var v19 int16 = 1000
    var v20 string = ("foo bar" + (("foo bar" + ("hello" + "foo bar")) + "a"))
    var v21 int8 = 1
    var v22 uint8 = v3
    v21 += 127
    for fi23 := 0; fi23 < 15; fi23++ {
        v22 = v22
        v2 = true
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}

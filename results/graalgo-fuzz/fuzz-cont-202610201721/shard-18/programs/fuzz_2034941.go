package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int, p1 int16) uint16 {
    if ((1e10 - 3.14159) >= 1e10) {
        v6 := []int{}
        var v7 float32 = 0.1
        _ = v6
        _ = v7
    } else {
        var v8 int64 = 1
        _ = v8
    }
    v9 := GenS1{F0: -84}
    if (1 != p1) {
        if false {
            v10 := []int{}
            p0 -= (p0 - 579)
            _ = v10
        } else {
            switch ("a" + "hello") {
            case "":
                if (true || false) {
                    v11 := []int{-895}
                    _ = v11
                } else {
                    var v12 float32 = 1.5
                    fmt.Println(p1, p1)
                    _ = v12
                }
            case "a":
                fmt.Println(p0, p1)
            case "hello":
                v13 := GenS1{F0: 95}
                _ = v13
            }
        }
        v14 := GenS1{F0: 1}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v9
    return 65535
}
func genFn1(p0 string, p1 int64, p2 uint) uint16 {
    var v15 int = 127
    var v16 int64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 32767
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 int64 = -653
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    for fi17 := 0; fi17 < 7; fi17++ {
        v1 = v1
        _ = fi17
    }
    v18 := GenS1{F0: 89}
    if (-155 >= 1) {
        fmt.Println(v1, v1)
    }
    if (!((genFn1(("foo bar" + ("\n\t" + "hello")), (v3 + (v3 + v3)), 42) > genFn1("hello", v3, 80)) || v1)) {
        v1 = ((!("hello" < "a")) || (false && (255 != 42)))
        fmt.Println(v3, v3)
    } else {
        var v19 string = (("\n\t" + "\n\t") + "a")
        v20 := GenS1{F0: 127}
        _ = v19
        _ = v20
    }
    fmt.Println(v1, v3)
    var v21 uint8 = 2
    var v22 uint = 65535
    fmt.Println(v3)
    v3 = v3
    if (536 == 65535) {
        var v23 int16 = 0
        switch 4294967295 {
        case 32768:
            var v24 float64 = ((3.14159 * ((3.14159 + 1.5) + ((0.0 / -2.25) - -2.25))) + 1e10)
            _ = v24
        case 2147483647:
            v25 := GenS1{F0: 127}
            _ = v25
        case 220:
            switch 1 {
            case 32768:
                v23 = (v23 & v23)
            default:
                var v26 uint32 = 1
                _ = v26
            }
        }
        _ = v23
    } else {
        v27 := []int{-9223372036854775808, -2147483648, -2147483648, 505}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}

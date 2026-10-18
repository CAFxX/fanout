package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16, p1 int16, p2 string) int {
    p2 = "a"
    if false {
        if false {
            var v6 uint = 42
            _ = v6
        }
        p2 = p2
        p2 = p2
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 825}
            _ = fi7
            _ = v8
        }
        fmt.Println(p1, p0, p1)
    }
    v9 := []int{128, 848, 127}
    for fi10 := 0; fi10 < 1; fi10++ {
        if false {
            if true {
                fmt.Println(p2)
            }
        } else {
            fmt.Println(fi10, p0)
        }
        if len(v9) > 0 {
            v9[0] = fi10
        }
        if true {
            for fi11 := 0; fi11 < 16; fi11++ {
                p2 = p2
                fmt.Println(p2, p0, p2)
                _ = fi11
            }
            var v12 int32 = 1
            fmt.Println(fi10, fi10)
            _ = v12
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 32767
}

func main() {
    var v1 float64 = 0.0
    var v2 bool = (false || ((true || false) && true))
    v3 := make(map[uint32]string)
    var v4 uint16 = 65535
    var v5 float64 = v1
    v13 := GenS1{F0: -721}
    v14 := GenS1{F0: -2147483648}
    switch "\n\t" {
    case "hello":
        v15 := []int{}
        _ = v15
    case "\n\t":
        v1 -= 1.5
    case "a":
        v2 = (492 != 4294967295)
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := v3[0]
        _ = fi16
        _ = v17
    }
    var v18 int8 = 127
    if (true || false) {
        v5 = (v5 / 3.14159)
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := []int{}
        var v21 int8 = 1
        if true {
            v3[42] = (("foo bar" + "") + "")
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 uint32 = 592
                _ = fi22
                _ = v23
            }
            if (!v2) {
                for fi24 := 0; fi24 < 7; fi24++ {
                    v18 -= (v21 >> ((42) & 63))
                    _ = fi24
                }
                v2 = (0 <= 0)
            }
        } else {
            var v25 uint32 = 42
            v25 = 278
            _ = v25
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v26 := GenS1{F0: -709}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}

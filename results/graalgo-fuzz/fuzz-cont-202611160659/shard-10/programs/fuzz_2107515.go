package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []int{127}
    var v2 uint = 0
    v3 := []string{"hello", (("a" + ("a" + ("\n\t" + "a"))) + "\n\t"), "hello"}
    var v4 int = 0
    var v5 string = "\n\t"
    v6 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 15; fi7++ {
        v5 = v5
        _ = fi7
    }
    v2 = 255
    fmt.Println(v5)
    v4 = (v4 * (v4 & 917))
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 uint32 = 0
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    if (3.14159 <= 1.5) {
        var v11 string = ((v5 + v5) + "hello")
        if (!(false && (float64(3.14159) < (-2.25 * (-2.25 * 1.5))))) {
            for fi12 := 0; fi12 < 16; fi12++ {
                v11 = v5
                for fi13 := 0; fi13 < 4; fi13++ {
                    fmt.Println(v4, v2)
                    delete(v6, 0)
                    _ = fi13
                }
                if (18446744073709551615 == 4294967295) {
                    var v14 uint = (v2 >> ((v2) & 63))
                    fmt.Println(v5)
                    var v15 int64 = -681
                    _ = v14
                    _ = v15
                }
                _ = fi12
            }
        }
        _ = v11
    }
    fmt.Println(v2)
    v2 += (^(v2))
    v16 := GenS1{F0: false}
    fmt.Println(v2)
    if false {
        v4 += 1
    }
    v5 = (((v5 + ((v5 + "") + "")) + v5) + ((v5 + v5) + ("foo bar" + v5)))
    var v17 uint32 = 65535
    var v18 uint32 = v17
    var v19 float64 = (3.14159 - -2.25)
    for fi20 := 0; fi20 < 3; fi20++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi20
    }
    v5 = "hello"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}

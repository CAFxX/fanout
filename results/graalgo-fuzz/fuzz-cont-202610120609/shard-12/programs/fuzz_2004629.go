package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int) uint64 {
    var v7 uint16 = 1000
    v7 -= 42
    v8 := []int{(p0 ^ (p0 + (p0 << ((4294967295) & 63))))}
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    v1 := []int{}
    var v2 uint16 = 255
    var v3 int16 = -490
    v4 := []string{""}
    var v5 uint64 = 0
    var v6 bool = (-2.25 > 0.0)
    v3 -= (v3 | v3)
    var v9 string = ("" + "hello")
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            v6 = ((v2 != 173) || (v2 > (^(v2))))
            var v12 bool = true
            _ = fi11
            _ = v12
        }
        var v13 float64 = 0.1
        var v14 uint32 = 2147483648
        _ = fi10
        _ = v13
        _ = v14
    }
    v6 = ((97 < -128) && ((1 == 814) || (42 <= 4294967295)))
    var v15 string = v9
    if true {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint = 2
            switch v9 {
            case "":
                if len(v4) > 0 {
                    v4[0] = v15
                }
            }
            _ = fi16
            _ = v17
        }
    } else {
        v18 := GenS1{F0: 130}
        _ = v18
    }
    var v19 bool = v6
    v15 = v9
    for fi20 := 0; fi20 < 18; fi20++ {
        v15 = "foo bar"
        _ = fi20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}

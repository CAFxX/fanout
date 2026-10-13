package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint8) uint8 {
    p0 -= (p0 << ((1) & 63))
    var v5 bool = true
    v6 := []int{32767, 4294967295, 605}
    p0 += p0
    var v7 string = ("foo bar" + ("" + ("\n\t" + (("a" + (("" + "a") + "")) + (("hello" + "a") + "\n\t")))))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (p0 % (((p0 | (p0 % ((p0) | 1)))) | 1))
}

func main() {
    v1 := []int{-776, 872, 2147483647, 4294967295}
    v2 := []int{900, -508, 515}
    var v3 bool = (!true)
    var v4 uint16 = 388
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = -32768
        v10 := GenS1{F0: 42}
        v9 = v9
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "foo bar"
    for fi12 := 0; fi12 < 15; fi12++ {
        v2 = append(v2, fi12)
        v13 := GenS1{F0: 1}
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v4, v3)
    var v14 int32 = 159
    for fi15 := 0; fi15 < 20; fi15++ {
        v11 = v11
        _ = fi15
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        if true {
            v17 := GenS1{F0: 9223372036854775807}
            v4 = 65535
            _ = v17
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(v4, fi18)
                var v19 uint = 255
                _ = fi18
                _ = v19
            }
        }
        _ = fi16
    }
    v20 := GenS1{F0: 1}
    var v21 float64 = ((1e10 + 0.0) - -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}

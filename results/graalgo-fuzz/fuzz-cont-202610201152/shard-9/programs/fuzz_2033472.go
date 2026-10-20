package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 int16, p1 uint32, p2 bool) int8 {
    switch -118 {
    case 127:
        if p2 {
            var v6 int64 = 889
            _ = v6
        }
    case 32768:
        if ((796 != 1) || p2) {
            p0 -= p0
            for fi7 := 0; fi7 < 1; fi7++ {
                fmt.Println(p2, p2)
                var v8 uint32 = (p1 % ((371) | 1))
                _ = fi7
                _ = v8
            }
            p0 = 1
        } else {
            if p2 {
                v9 := GenS1{F0: 737, F1: 1}
                v10 := GenS1{F0: 2, F1: 42}
                if true {
                    var v11 uint64 = 1
                    _ = v11
                }
                _ = v10
                _ = v9
            }
        }
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(p2)
        if true {
            for fi13 := 0; fi13 < 15; fi13++ {
                p2 = true
                v14 := []int{fi12, 211, fi12, (fi13 + (fi13 - 2147483647))}
                _ = fi13
                _ = v14
            }
        }
        p0 = p0
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int64 = 42
    var v2 int64 = v1
    v3 := []string{(("" + "hello") + "foo bar")}
    var v4 float64 = 0.0
    var v5 uint8 = 127
    v4 += (((0.1 + v4) * (1.5 - ((v4 + (1.5 / float64(1.5))) + float64(0.1)))) - float64(0.0))
    v2 += v1
    var v15 int8 = 100
    var v16 bool = (525 <= 4294967295)
    var v17 bool = false
    fmt.Println(v1)
    var v18 string = ""
    if len(v3) > 0 {
        v3[0] = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        _ = len(v3)
        v20 := GenS1{F0: 0, F1: 325}
        for fi21 := 0; fi21 < 4; fi21++ {
            if false {
                fmt.Println(fi19)
                var v22 bool = (v17 || v17)
                v23 := GenS1{F0: 940, F1: 2147483647}
                var v24 float64 = (v4 + v4)
                var v25 int16 = int16(fi21)
                _ = v22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    _ = len(v3)
    fmt.Println(v5, v17, v16)
    var v26 string = "a"
    v1 += 1
    var v27 float32 = 1e10
    if len(v3) > 0 {
        v3[0] = (v18 + (((("\n\t" + "\n\t") + v26) + ("\n\t" + v18)) + v26))
    }
    v1 = (v2 >> ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}

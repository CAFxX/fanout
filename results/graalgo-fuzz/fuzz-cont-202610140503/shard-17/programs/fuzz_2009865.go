package main

import (
    "fmt"
)


func main() {
    var v1 float64 = float64(1e10)
    v2 := []uint32{18, 2, 2147483647}
    var v3 float32 = (0.1 * -2.25)
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    var v6 bool = (255 == 127)
    v6 = (((!(127 < 1)) || ((45 >= 568) && (!false))) && v6)
    v1 = v1
    var v7 bool = true
    v1 += 1e10
    v2 = append(v2, 42)
    var v8 int32 = 367
    var v9 float32 = float32(v1)
    if (v8 >= v8) {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 string = ("\n\t" + "")
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                v6 = (v6 || true)
                var v14 float32 = (v3 * v3)
                var v15 float32 = v9
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
        fmt.Println(v3)
    } else {
        for fi16 := 0; fi16 < 7; fi16++ {
            fmt.Println(v1, v6)
            switch (("a" + "") + ("foo bar" + "foo bar")) {
            case "foo bar":
                v9 = v3
            case "a":
                for fi17 := 0; fi17 < 16; fi17++ {
                    v2 = append(v2, 548)
                    v3 -= ((3.14159 / 1.0) - (v9 * 0.0))
                    _ = fi17
                }
            case "hello":
                var v18 uint8 = 2
                _ = v18
            }
            v5[65535] = (fi16 - fi16)
            _ = fi16
        }
    }
    fmt.Println(v1)
    var v19 float64 = v1
    v9 = v9
    v19 = -2.25
    v19 = 0.1
    delete(v4, 714)
    var v20 uint = 4294967295
    _ = len(v2)
    v1 += v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []int{32768, 132}
    var v2 float32 = (1.5 * (3.14159 + -2.25))
    var v3 uint16 = 42
    var v4 int32 = 0
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 32767}
        var v7 uint64 = 0
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = false
    _ = len(v1)
    if (("" == ("hello" + ("hello" + ("\n\t" + "\n\t")))) || (!false)) {
        var v9 int32 = v4
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if (127 < 127) {
            if v8 {
                if ((65535 > 0) || v8) {
                    var v11 float64 = 1.5
                    v1 = append(v1, fi10)
                    for fi12 := 0; fi12 < 14; fi12++ {
                        v13 := GenS2{F0: 65535}
                        var v14 int32 = (v4 + -845)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
            }
            fmt.Println(v8, v3, v3)
            v15 := GenS2{F0: 1}
            _ = v15
        }
        var v16 uint16 = (v3 & (v3 >> ((731) & 63)))
        var v17 uint64 = 42
        _ = fi10
        _ = v16
        _ = v17
    }
    fmt.Println(v8)
    var v18 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}

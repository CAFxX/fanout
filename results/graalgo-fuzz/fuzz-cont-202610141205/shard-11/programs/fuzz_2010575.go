package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int32 = 42
    v2 := []int{42}
    var v3 int = 4294967295
    for fi4 := 0; fi4 < 16; fi4++ {
        v1 -= (^(v1))
        _ = fi4
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 := GenS1{F0: 127}
    v6 := GenS1{F0: 42}
    var v7 uint = 2
    v1 -= v1
    var v8 uint64 = 415
    v3 = (v3 - -470)
    var v9 float32 = (((float32(-2.25) - 0.1) * 3.14159) + 1e10)
    switch v3 {
    case -9223372036854775808:
        switch 4294967295 {
        case 32768:
            for fi10 := 0; fi10 < 17; fi10++ {
                fmt.Println(v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi10
            }
        case 124:
            fmt.Println(v3, v3, v7)
        }
    case 0:
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 float32 = 3.14159
                if len(v2) > 0 {
                    v2[0] = (fi11 << ((686) & 63))
                }
                v1 = v1
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                _ = len(v2)
                v15 := GenS1{F0: -128}
                _ = fi14
                _ = v15
            }
            _ = fi11
        }
    case 9223372036854775807:
        switch (v3 / ((v3) | 1)) {
        case 823:
            fmt.Println(v7)
        }
    }
    var v16 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}

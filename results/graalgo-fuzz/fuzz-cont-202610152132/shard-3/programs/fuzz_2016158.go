package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []uint32{1, 2147483647, 2147483648}
    v2 := []int{}
    var v3 int8 = 127
    var v4 int16 = 0
    _ = len(v1)
    v3 -= v3
    v3 = 1
    v4 = 32767
    v5 := GenS1{F0: 127}
    v3 = v3
    fmt.Println(v4, v4)
    if true {
        v4 = (^(v4))
        switch 128 {
        case 1:
            _ = len(v1)
        default:
            v3 = 100
        }
        for fi6 := 0; fi6 < 17; fi6++ {
            if true {
                v7 := GenS1{F0: 128}
                fmt.Println(v3)
                _ = v7
            }
            var v8 bool = ("\n\t" == "\n\t")
            var v9 uint32 = 1
            _ = fi6
            _ = v8
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            if false {
                v11 := GenS1{F0: 42}
                _ = v11
            }
            v2 = append(v2, (fi10 << ((531) & 63)))
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 int64 = 1
                var v14 int16 = v4
                if false {
                    if false {
                        var v15 bool = false
                        if (v14 <= -881) {
                            fmt.Println(fi10, v4, v4)
                        } else {
                            var v16 int16 = 42
                            _ = v16
                        }
                        switch "a" {
                        case "a":
                            v17 := GenS1{F0: 2}
                            _ = v17
                        }
                        _ = v15
                    } else {
                        v1 = append(v1, 0)
                        var v18 int = int(v13)
                        v4 -= v4
                        _ = v18
                    }
                } else {
                    var v19 uint = 65535
                    switch "hello" {
                    case "hello":
                        v2 = append(v2, fi10)
                    default:
                        v3 -= v3
                    }
                    _ = v19
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi10
        }
        v20 := GenS1{F0: 82}
        _ = v20
    }
    fmt.Println(v3, v4)
    v21 := GenS2{F0: 760}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}

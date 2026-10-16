package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0() int {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 int32 = 2147483647
        v6 := GenS1{F0: 18446744073709551615}
        var v7 uint32 = 42
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 0}
    _ = v8
    return 30
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 2
    var v3 uint8 = 1
    var v9 uint = (v2 + (v2 % ((42) | 1)))
    var v10 int64 = -9223372036854775808
    for fi11 := 0; fi11 < 12; fi11++ {
        switch ("\n\t" + "foo bar") {
        case "a":
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint32 = 295
                fmt.Println(fi12)
                _ = fi12
                _ = v13
            }
        }
        _ = fi11
    }
    var v14 bool = (374 >= 1)
    switch 128 {
    case 1:
        switch 2147483647 {
        case 32767:
            if (!true) {
                v15 := []int{}
                v16 := GenS1{F0: 2}
                _ = v15
                _ = v16
            }
        case 32768:
            var v17 uint8 = v3
            _ = v17
        }
    default:
        if (2 == 65535) {
            var v18 int64 = v10
            if true {
                var v19 bool = v14
                switch "hello" {
                case "hello":
                    if ((v9 / ((939) | 1)) <= v9) {
                        switch 32767 {
                        case 42:
                            v20 := []int{537, -2147483648, 4294967295, genFn0()}
                            _ = v20
                        case 128:
                            v21 := GenS1{F0: 4294967295}
                            _ = v21
                        case -490:
                            v22 := GenS1{F0: 0}
                            _ = v22
                        default:
                            v10 = 190
                        }
                        for fi23 := 0; fi23 < 19; fi23++ {
                            v9 += v2
                            var v24 uint8 = v3
                            _ = fi23
                            _ = v24
                        }
                        var v25 uint16 = 255
                        _ = v25
                    }
                case "":
                    v26 := GenS1{F0: 1}
                    _ = v26
                }
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 int8 = 127
                    v29 := []int{32768, fi27}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v19
            }
            v30 := []int{366, 244, 32767, -61}
            _ = v18
            _ = v30
        }
    }
    v2 = (v9 / ((uint(v3)) | 1))
    v31 := GenS2{F0: 4294967295}
    v32 := GenS2{F0: 65535}
    var v33 uint = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v9
}

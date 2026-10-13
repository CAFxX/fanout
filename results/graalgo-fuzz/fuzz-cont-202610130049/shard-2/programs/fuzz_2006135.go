package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() int32 {
    v4 := GenS1{F0: 127}
    var v5 uint16 = 32767
    v5 += v5
    _ = v4
    _ = v5
    return 0
}
func genFn1() uint8 {
    v6 := []int{582, 42}
    var v7 float64 = (1e10 / (0.1 - (0.0 * ((1e10 * 0.1) * float64((3.14159 - (1.5 / float32(3.14159))))))))
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int32 = 53
        for fi10 := 0; fi10 < 2; fi10++ {
            v7 = v7
            if len(v6) > 0 {
                v6[0] = 32768
            }
            var v11 int32 = (v9 >> ((255) & 63))
            _ = fi10
            _ = v11
        }
        v7 = v7
        _ = fi8
        _ = v9
    }
    switch 0 {
    case 32768:
        if (-2147483648 < 0) {
            var v12 uint = 0
            for fi13 := 0; fi13 < 6; fi13++ {
                v12 = v12
                var v14 uint64 = 4294967295
                var v15 float64 = v7
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        } else {
            v7 = 1e10
            for fi16 := 0; fi16 < 8; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    v18 := GenS1{F0: 42}
                    v7 = float64(-2.25)
                    var v19 int = 763
                    var v20 int64 = 773
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                switch ("\n\t" + "\n\t") {
                case "":
                    v21 := GenS1{F0: 0}
                    _ = v21
                default:
                    v22 := GenS1{F0: 1}
                    _ = v22
                }
                var v23 uint = 255
                _ = fi16
                _ = v23
            }
        }
    default:
        var v24 int16 = 42
        _ = v24
    }
    var v25 int32 = 42
    _ = v25
    _ = v6
    _ = v7
    return 128
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 1
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint32 = 392
        _ = fi26
        _ = v27
    }
    if true {
        var v28 int16 = -430
        _ = v28
    }
    v29 := v2[2147483647]
    var v30 uint8 = 0
    var v31 uint8 = 1
    if false {
        v30 = v31
        v1[-464] = ((v30 / ((v30) | 1)) > v30)
    } else {
        v32 := []int{}
        var v33 int32 = -867
        v2[1] = 1
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: (v30 - (v31 * v31))}
    var v35 int = -110
    var v36 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
}

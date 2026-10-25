package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0() uint8 {
    var v4 uint16 = 255
    v5 := GenS1{F0: 4294967295}
    _ = v4
    _ = v5
    return 2
}
func genFn1() uint32 {
    var v6 bool = (!(false && (!(((-2.25 + 1e10) <= 1.5) || false))))
    var v7 int32 = 1
    v6 = false
    var v8 bool = ((0.1 * 0.1) <= -2.25)
    var v9 uint64 = 0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 412
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]bool)
    var v3 int = -135
    v10 := GenS1{F0: 853}
    v11 := []int{v3, (v3 / (((v3 >> ((1) & 63))) | 1)), (v3 % ((298) | 1))}
    var v12 uint32 = 2
    for fi13 := 0; fi13 < 16; fi13++ {
        switch fi13 {
        case 2147483647:
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        var v14 bool = (!true)
        v15 := GenS2{F0: 18446744073709551615}
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            v12 = 2
            _ = fi17
        }
        _ = fi16
    }
    v11 = append(v11, (v3 % (((v3 * (v3 | 54))) | 1)))
    if false {
        for fi18 := 0; fi18 < 13; fi18++ {
            v12 += 2147483648
            for fi19 := 0; fi19 < 2; fi19++ {
                if len(v11) > 0 {
                    v11[0] = (fi18 << ((65535) & 63))
                }
                delete(v2, "\n\t")
                _ = fi19
            }
            _ = fi18
        }
        var v20 int = (v3 & -934)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint8 = 2
            _ = len(v11)
            var v23 uint8 = (v22 + 92)
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    var v24 int64 = 104
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
}

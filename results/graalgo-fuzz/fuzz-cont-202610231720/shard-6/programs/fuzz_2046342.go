package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0() int8 {
    var v4 string = ("\n\t" + "a")
    for fi5 := 0; fi5 < 2; fi5++ {
        if (false && (13 == 1)) {
            fmt.Println(fi5)
        }
        v6 := GenS1{F0: 1, F1: 121, F2: 128}
        if true {
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: 0, F1: 2, F2: 209}
                fmt.Println(fi7, fi5, v4)
                _ = fi7
                _ = v8
            }
            if false {
                v9 := GenS1{F0: 65535, F1: 42, F2: 0}
                v4 = v4
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := []int{}
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 uint16 = 2
                    var v14 int16 = 32767
                    fmt.Println(fi5, fi5, v14)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := []int{fi15, fi5}
                    if false {
                        v4 = v4
                    }
                    fmt.Println(fi5, v4, v4)
                    _ = fi15
                    _ = v16
                }
                var v17 bool = (42 == 255)
                _ = v17
            }
            if ((100 >= 0) && false) {
                for fi18 := 0; fi18 < 17; fi18++ {
                    v4 = ((v4 + "foo bar") + v4)
                    switch 4294967295 {
                    case 128:
                        v4 = (v4 + v4)
                    default:
                        v4 = v4
                    }
                    _ = fi18
                }
                v4 = v4
            }
        } else {
            var v19 int = (fi5 / ((fi5) | 1))
            _ = v19
        }
        _ = fi5
        _ = v6
    }
    _ = v4
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v20 int = -802
    if false {
        v2[65535] = 2147483648
    }
    v20 = v20
    v20 = -2147483648
    v21 := []int{v20, (v20 << ((42) & 63))}
    v20 -= (v20 % (((^(v20))) | 1))
    delete(v3, "")
    v20 -= 0
    if true {
        if (0 != -72) {
            fmt.Println(v20)
            fmt.Println(v20)
            var v22 bool = (-726 != 532)
            _ = v22
        } else {
            delete(v1, 2)
            v21 = append(v21, v20)
        }
        delete(v2, 2147483647)
    } else {
        v20 += -421
    }
    v23 := GenS1{F0: 255, F1: 1, F2: 62}
    v2[4294967295] = 914
    for fi24 := 0; fi24 < 4; fi24++ {
        v20 = (v20 + -966)
        var v25 string = (("a" + "hello") + "\n\t")
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 0, F1: 214, F2: 3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v26
    _ = v3
}

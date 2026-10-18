package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint8) uint64 {
    p0 -= (^(p0))
    if (!(!(42 <= 42))) {
        fmt.Println(p1, p1)
        p1 = p1
        for fi4 := 0; fi4 < 4; fi4++ {
            switch fi4 {
            case -554:
                for fi5 := 0; fi5 < 18; fi5++ {
                    var v6 int8 = 127
                    if ((fi5 - (-(fi5))) <= fi5) {
                        v7 := []int{32768, -311, fi4}
                        p0 += -1
                        v7 = append(v7, 2147483647)
                        _ = v7
                    } else {
                        var v8 float64 = 0.1
                        var v9 bool = ((fi4 | fi4) == (fi4 - -1))
                        _ = v8
                        _ = v9
                    }
                    _ = fi5
                    _ = v6
                }
            case -291:
                v10 := []int{fi4, fi4, fi4}
                _ = v10
            case 254:
                var v11 uint8 = p1
                _ = v11
            }
            if false {
                v12 := []int{fi4, (^(fi4)), (fi4 - fi4)}
                _ = v12
            } else {
                v13 := GenS1{F0: false, F1: int64(fi4)}
                v14 := []int{32768, fi4, fi4, (-(fi4))}
                _ = len(v14)
                _ = v13
                _ = v14
            }
            _ = fi4
        }
    }
    var v15 uint32 = 2
    _ = p0
    _ = p1
    _ = v15
    return 0
}

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]int)
    var v3 int8 = (v1 * (^(v1)))
    fmt.Println(v1)
    v16 := GenS2{F0: 65535, F1: 42}
    var v17 string = "foo bar"
    var v18 float64 = 0.0
    v3 += (v3 * v3)
    switch 2147483647 {
    case 32767:
        for fi19 := 0; fi19 < 16; fi19++ {
            delete(v2, 2147483647)
            _ = fi19
        }
    case 1:
        switch 32768 {
        case 974:
            v17 = ("foo bar" + v17)
        }
    }
    var v20 uint16 = 0
    v21 := GenS3{F0: (v20 % (((v20 & v20)) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
